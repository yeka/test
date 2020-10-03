package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	//log.Println(CreateCA("go-ca"))
	//caKey, caCert, caErr := LoadCA("go-ca")
	//if caErr != nil {
	//	log.Println(caErr)
	//	return
	//}
	//_, _ = caKey, caCert
	//log.Println(CreateCert(caKey, caCert, "go-cert"))
	//return

	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("This is an example server.\n"))
	})

	//if err := NoCASelfSigned(handler, ":443"); err != nil {
	//	log.Println(err)
	//}

	//caKey, caCert, caErr := LoadCA("go-ca")
	//if caErr != nil {
	//	log.Println(caErr)
	//	return
	//}
	//if err := CA2(caKey, caCert, handler, ":443"); err != nil {
	//	log.Println(err)
	//}


	//err := filebasedCertificate(x)

	//if err := CA(x); err != nil {
	//	log.Println(err)
	//}

	s := http.Server{Addr: ":443", Handler: handler}
	if err := s.ListenAndServeTLS("https-server.crt", "https-server.key"); err != nil {
		log.Println(err)
	}
}

func filebasedCertificate(handler http.Handler) error {
	s := http.Server{Addr: ":1443", Handler: handler}
	return s.ListenAndServeTLS("https-server.crt", "https-server.key")
}

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}

// NoCASelfSigned
// reference: https://gist.github.com/samuel/8b500ddd3f6118d052b5e6bc16bc4c09
func NoCASelfSigned(handler http.Handler, addr string) error {
	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}
	template := x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{"Personal Company, INC."},
			Country:      []string{"ID"},
			Locality: []string{"Depok"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().AddDate(1, 0, 0),

		//KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		//ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
		//BasicConstraintsValid: true,

		DNSNames: []string{"personal.dev"},
	}

	/*
	   hosts := strings.Split(*host, ",")
	   for _, h := range hosts {
	   	if ip := net.ParseIP(h); ip != nil {
	   		template.IPAddresses = append(template.IPAddresses, ip)
	   	} else {
	   		template.DNSNames = append(template.DNSNames, h)
	   	}
	   }
	   if *isCA {
	   	template.IsCA = true
	   	template.KeyUsage |= x509.KeyUsageCertSign
	   }
	*/

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &certPrivKey.PublicKey, certPrivKey)
	if err != nil {
		log.Fatalf("Failed to create certificate: %s", err)
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	certPrivKeyPEM := pem.EncodeToMemory(pemBlockForKey(certPrivKey))

	serverCert, err := tls.X509KeyPair(certPEM, certPrivKeyPEM)
	if err != nil {
		return err
	}

	serverTLSConf := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
	}

	// for Client
	//certpool := x509.NewCertPool()
	//certpool.AppendCertsFromPEM(caPEM.Bytes())
	//clientTLSConf := &tls.Config{
	//	RootCAs: certpool,
	//}

	srv := http.Server{Addr: addr, Handler: handler, TLSConfig: serverTLSConf}
	return srv.ListenAndServeTLS("", "")
}

func CA(handler http.Handler, addr string) error {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization: []string{"Personal CA, INC."},
			Country:      []string{"ID"},
			//Province:      []string{""},
			Locality: []string{"Jakarta"},
			//StreetAddress: []string{""},
			//PostalCode:    []string{""},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return err
	}

	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})

	return CA2(caPrivKey, ca, handler, addr)
}

func CA2(caKey *rsa.PrivateKey, caCert *x509.Certificate, handler http.Handler, addr string) error {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{"Personal Company, INC."},
			Country:      []string{"ID"},
			//Province:      []string{""},
			Locality: []string{"Depok"},
			//StreetAddress: []string{"Golden Gate Bridge"},
			//PostalCode:    []string{"94016"},
		},
		//IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		DNSNames:     []string{"personal.dev", "two.personal.dev", "*.nekochan.dev", "nekochan.dev"},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(2, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, caCert, &certPrivKey.PublicKey, caKey)
	if err != nil {
		return err
	}

	certPEM := new(bytes.Buffer)
	pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	certPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certPrivKey),
	})

	serverCert, err := tls.X509KeyPair(certPEM.Bytes(), certPrivKeyPEM.Bytes())
	if err != nil {
		return err
	}

	serverTLSConf := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
	}

	// for Client
	//certpool := x509.NewCertPool()
	//certpool.AppendCertsFromPEM(caPEM.Bytes())
	//clientTLSConf := &tls.Config{
	//	RootCAs: certpool,
	//}

	srv := http.Server{Addr: addr, Handler: handler, TLSConfig: serverTLSConf}
	return srv.ListenAndServeTLS("", "")
}

func CreateCA(filename string) error {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization: []string{"Personal CA, INC."},
			Country:      []string{"ID"},
			//Province:      []string{""},
			Locality: []string{"Jakarta"},
			//StreetAddress: []string{""},
			//PostalCode:    []string{""},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	caCertBytes, err := x509.CreateCertificate(rand.Reader, ca, ca, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return err
	}

	caPEM := new(bytes.Buffer)
	pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caCertBytes,
	})

	caPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})

	if err := ioutil.WriteFile(filename+"-key.pem", caPrivKeyPEM.Bytes(), 0655); err != nil {
		return err
	}

	return ioutil.WriteFile(filename+"-crt.pem", caPEM.Bytes(), 0655)
}

func LoadCA(filename string) (*rsa.PrivateKey, *x509.Certificate, error) {
	privPEMBytes, err := ioutil.ReadFile(filename + "-key.pem")
	if err != nil {
		return nil, nil, err
	}

	privBytes, _ := pem.Decode(privPEMBytes)
	if privBytes.Type != "RSA PRIVATE KEY" {
		return nil, nil, errors.New("RSA private key is of the wrong type: " + privBytes.Type)
	}

	caKey, err := x509.ParsePKCS1PrivateKey(privBytes.Bytes)
	if err != nil {
		return nil, nil, err
	}

	certPEMBytes, err := ioutil.ReadFile(filename + "-crt.pem")
	if err != nil {
		return nil, nil, err
	}

	certBytes, _ := pem.Decode(certPEMBytes)
	if certBytes.Type != "CERTIFICATE" {
		return nil, nil, errors.New("Certificate is of the wrong type: " + certBytes.Type)
	}
	caCert, err := x509.ParseCertificate(certBytes.Bytes)
	if err != nil {
		return nil, nil, err
	}

	return caKey, caCert, nil
}

func CreateCert(caKey *rsa.PrivateKey, caCert *x509.Certificate, filename string) error {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1658),
		Subject: pkix.Name{
			Organization: []string{"Personal Company, INC."},
			Country:      []string{"ID"},
			//Province:      []string{""},
			Locality: []string{"Depok"},
			//StreetAddress: []string{"Golden Gate Bridge"},
			//PostalCode:    []string{"94016"},
		},
		IPAddresses:  []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
		NotBefore:    time.Now(),
		NotAfter:     time.Now().AddDate(10, 0, 0),
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}

	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, caCert, &certPrivKey.PublicKey, caKey)
	if err != nil {
		return err
	}

	certPEM := new(bytes.Buffer)
	pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	certPrivKeyPEM := new(bytes.Buffer)
	pem.Encode(certPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(certPrivKey),
	})

	//serverCert, err := tls.X509KeyPair(certPEM.Bytes(), certPrivKeyPEM.Bytes())
	//if err != nil {
	//	return err
	//}

	if err := ioutil.WriteFile(filename+"-key.pem", certPrivKeyPEM.Bytes(), 0655); err != nil {
		return err
	}
	if err := ioutil.WriteFile(filename+"-crt.pem", certPEM.Bytes(), 0655); err != nil {
		return err
	}

	return nil
}
