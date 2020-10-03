package libs

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/ssh"
)

const (
	ErrorLevel int = 1
	DebugLevel int = 2
)

type TunnelConnection struct {
	localPort string
	listener  net.Listener

	closing chan struct{} // to signal all opened connection to close
	wg      sync.WaitGroup
}

// Tunnel provides ssh tunnel connections
// Params:
// - sshDSN   ssh server using DSN format "user:pass@server:port" (port default to 22 if not provided)
// - remote   remote address "ip:port"
// - local    optional local address ":port" (when local not provided, random port will be used)
//            opened local port can be requested using *TunnelConnection.LocalPort()
// - authMethods an optional additional authentication method. some helper for this options
//               SSHPassword(string) - password authentication (automatically used when password is present in DSN)
//               SSHKey(io.Reader) - SSH Key authentication, accept io.Reader so file or string can be used
//
// Example of simple usage:
//    Tunnel("user:password@server", "redis_server:6379", "")
//
// Example of using string based SSH Key
//    Tunnel("user@server", "redis_server:6379", "", SSHKey(strings.NewReader(sshkey)))
//
func Tunnel(sshDSN, remote, local string, authMethods ...ssh.AuthMethod) (*TunnelConnection, error) {
	// Validate input
	user, pass, host, port, err := parseDsn(sshDSN)
	if err != nil {
		return nil, fmt.Errorf("invalid SSH server: %w", err)
	}
	if user == "" || host == "" {
		return nil, errors.New("invalid SSH server address: user and host must be provided")
	}
	if port == "" {
		port = "22"
	}

	_, _, rhost, rport, err := parseDsn(remote)
	if err != nil {
		return nil, fmt.Errorf("invalid remote address: %w", err)
	}
	if rhost == "" || rport == "" {
		return nil, errors.New("invalid remote address: host and port must be provided")
	}

	if _, _, _, _, err := parseDsn(local); err != nil {
		return nil, fmt.Errorf("invalid local address: %w", err)
	}

	// Preparing configuration
	cfg := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{SSHPassword(pass)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	for _, v := range authMethods {
		if v == nil {
			continue
		}
		cfg.Auth = append(cfg.Auth, v)
	}

	// Start listening on local port
	listener, err := net.Listen("tcp", local)
	if err != nil {
		return nil, fmt.Errorf("unable to open local port (%v): %w", local, err)
	}

	tc := &TunnelConnection{
		strconv.Itoa(listener.Addr().(*net.TCPAddr).Port),
		listener,
		make(chan struct{}),
		sync.WaitGroup{},
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				if strings.Contains(err.Error(), "use of closed network connection") {
					tc.log(DebugLevel, "Tunnel closed")
					return
				} else {
					tc.log(ErrorLevel, fmt.Sprintf("accepting connection error: %#v\n", err))
				}
				return
			}

			tc.log(DebugLevel, "new connection %v => %v", conn.LocalAddr(), conn.RemoteAddr())

			go func() {
				tc.wg.Add(1)
				defer tc.wg.Done()
				tc.spawnConnection(conn, host+":"+port, cfg, remote)
			}()
		}
	}()
	return tc, nil
}

// LocalPort show local port opened for local (tunnelled) connection
func (tc *TunnelConnection) LocalPort() string {
	return tc.localPort
}

// Close will close the tunnel
func (tc *TunnelConnection) Close() {
	close(tc.closing)
	tc.wg.Wait()

	err := tc.listener.Close()
	if err != nil {
		err = fmt.Errorf("closing listener: %w", err)
		tc.log(ErrorLevel, err.Error())
	}
}

func (tc *TunnelConnection) spawnConnection(localConn net.Conn, server string, cfg *ssh.ClientConfig, remote string) {
	defer tc.close(localConn, "closing local connection")

	serverConn, err := ssh.Dial("tcp", server, cfg)
	if err != nil {
		err = fmt.Errorf("server dial error: %w", err)
		tc.log(ErrorLevel, err.Error())
		return
	}
	defer tc.close(serverConn, "closing server connection")

	remoteConn, err := serverConn.Dial("tcp", remote)
	if err != nil {
		err = fmt.Errorf("remote dial error: %w", err)
		tc.log(ErrorLevel, err.Error())
		return
	}
	defer tc.close(remoteConn, "closing remote connection")

	copyConn := func(writer, reader net.Conn) {
		_, err := io.Copy(writer, reader)
		if err != nil {
			err = fmt.Errorf("io.Copy error: %w", err)
			tc.log(ErrorLevel, err.Error())
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)

	<-tc.closing
	return
}

func (tc *TunnelConnection) close(c io.Closer, onErrorMessage string) {
	err := c.Close()
	if err != nil {
		err = fmt.Errorf(onErrorMessage+": %w", err)
		tc.log(ErrorLevel, onErrorMessage)
	}
}

func (tc *TunnelConnection) log(level int, format string, a ...interface{}) {
	severity := map[int]string{
		ErrorLevel: "ERROR",
		DebugLevel: "DEBUG",
	}[level]
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), severity, fmt.Sprintf(format, a...))
}

/*
	======================
	Authentication Methods
	======================
*/

func SSHKey(r io.Reader) ssh.AuthMethod {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(b)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}

func SSHPassword(pass string) ssh.AuthMethod {
	return ssh.Password(pass)
}

/*
	=================
	DSN related tools
	=================
*/

var dsnRegex = regexp.MustCompile("^(([a-zA-Z0-9_.]+?)(:[a-zA-Z0-9_.]+?)?@)?([a-zA-Z0-9_.]+?)?(:[0-9_.]+?)?$")

func parseDsn(s string) (user, pass, host, port string, err error) {
	x := dsnRegex.FindAllStringSubmatch(s, 1)
	if len(x) != 1 {
		err = errors.New("invalid DSN:" + s)
		return
	}
	user = x[0][2]
	pass = strings.TrimLeft(x[0][3], ":")
	host = strings.TrimLeft(x[0][4], "@")
	port = strings.TrimLeft(x[0][5], ":")
	return
}

func reconstructDSN(user, pass, host, port string) string {
	if pass != "" {
		user += ":" + pass
	}
	if port != "" {
		host += ":" + port
	}
	if user != "" && host != "" {
		return user + "@" + host
	}
	return user + host
}
