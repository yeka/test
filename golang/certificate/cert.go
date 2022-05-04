package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"
)

var rootCert = x509.Certificate{
	Subject: pkix.Name{
		Country:      []string{"ID"},
		Organization: []string{"YK Brothers Co."},
		CommonName:   "Yeka Root CA",
	},
	NotBefore:   time.Now(),
	NotAfter:    time.Now().AddDate(10, 0, 0),
	IPAddresses: []net.IP{},
	DNSNames:    []string{},
}

var intermediateCert = x509.Certificate{
	Subject: pkix.Name{
		Country:      []string{"ID"},
		Organization: []string{"YK Intermediate Corp"},
		CommonName:   "Yeka Intermediate CA",
	},
	NotBefore:   time.Now(),
	NotAfter:    time.Now().AddDate(10, 0, 0),
	IPAddresses: []net.IP{},
	DNSNames:    []string{},
}

var serverCert = x509.Certificate{
	Subject: pkix.Name{
		Country:      []string{"ID"},
		Organization: []string{"Go Web Corp."},
		CommonName:   "Go Web",
	},
	NotBefore:   time.Now(),
	NotAfter:    time.Now().AddDate(10, 0, 0),
	IPAddresses: []net.IP{net.ParseIP("127.0.0.1"), net.ParseIP("192.168.1.6")},
	DNSNames:    []string{},
}

// In this example you'll find Cert1, Cert2 & Cert3
// Cert1 creates a self signed certificate
// Cert2 creates a self signed root certificate authority & a server certificate
// Cert3 creates a self signed root certificate authority, an intermediate certificate & a server certificate

func main() {
	fmt.Println(
		// Cert1(ECDSA, "cert.pem", "key.pem"),
		Cert2(ECDSA, "root_cert.pem", "root_key.pem", "cert.pem", "key.pem"),
		// Cert3(ECDSA, "root_cert.pem", "root_key.pem", "intermediate_cert.pem", "intermediate_key.pem", "cert.pem", "key.pem"),
	)
}

// Cert1 creates a self signed certificate
func Cert1(generateFn PrivateKeyGenerator, certFilename, keyFilename string) error {
	key, err := generateFn()
	if err != nil {
		return err
	}

	cert := ServerCert(serverCert)

	b, err := x509.CreateCertificate(rand.Reader, &cert, &cert, key.Public(), key)
	if err != nil {
		return fmt.Errorf("Failed to create certificate: %s", err)
	}

	return writeToFile(b, key, certFilename, keyFilename)
}

// Cert2 creates a self signed root certificate authority & a server certificate
func Cert2(generateFn PrivateKeyGenerator, rootCertFilename, rootKeyFilename, certFilename, keyFilename string) error {
	// Root Certificate
	var rCert *x509.Certificate
	var rKey PrivateKey
	var err error

	rCert, rKey, err = loadFromFile(rootCertFilename, rootKeyFilename)
	if err != nil {
		rKey, err = generateFn()
		if err != nil {
			return err
		}

		cert := RootCert(rootCert)
		rCert = &cert
		rootBytes, err := x509.CreateCertificate(rand.Reader, rCert, rCert, rKey.Public(), rKey)
		if err != nil {
			return fmt.Errorf("Failed to create certificate: %s", err)
		}
		if err := writeToFile(rootBytes, rKey, rootCertFilename, rootKeyFilename); err != nil {
			return err
		}
	}

	// Server Certificate
	key, err := generateFn()
	if err != nil {
		return err
	}

	cert := ServerCert(serverCert)

	b, err := x509.CreateCertificate(rand.Reader, &cert, rCert, key.Public(), rKey)
	if err != nil {
		return fmt.Errorf("Failed to create certificate: %s", err)
	}

	return writeToFile(b, key, certFilename, keyFilename)
}

// Cert3 creates a self signed root certificate authority, an intermediate certificate & a server certificate
func Cert3(generateFn PrivateKeyGenerator, rootCertFilename, rootKeyFilename, intermediateCertFilename, intermediateKeyFilename, certFilename, keyFilename string) error {
	// Root Certificate
	var rCert *x509.Certificate
	var rKey PrivateKey
	var err error

	rCert, rKey, err = loadFromFile(rootCertFilename, rootKeyFilename)
	if err != nil {
		rKey, err = generateFn()
		if err != nil {
			return err
		}

		c := RootCert(rootCert)
		rCert = &c
		rBytes, err := x509.CreateCertificate(rand.Reader, rCert, rCert, rKey.Public(), rKey)
		if err != nil {
			return fmt.Errorf("Failed to create certificate: %s", err)
		}
		if err := writeToFile(rBytes, rKey, rootCertFilename, rootKeyFilename); err != nil {
			return err
		}
	}

	// Interediate Certificate
	var iCert *x509.Certificate
	var iKey PrivateKey

	iCert, iKey, err = loadFromFile(intermediateCertFilename, intermediateKeyFilename)
	if err != nil {
		rKey, err = generateFn()
		if err != nil {
			return err
		}

		c := IntermediateCert(intermediateCert)
		iCert = &c
		iBytes, err := x509.CreateCertificate(rand.Reader, iCert, rCert, iKey.Public(), rKey)
		if err != nil {
			return fmt.Errorf("Failed to create certificate: %s", err)
		}
		if err := writeToFile(iBytes, iKey, intermediateCertFilename, intermediateKeyFilename); err != nil {
			return err
		}
	}

	// Server Certificate
	key, err := generateFn()
	if err != nil {
		return err
	}

	cert := ServerCert(serverCert)

	b, err := x509.CreateCertificate(rand.Reader, &cert, iCert, key.Public(), iKey)
	if err != nil {
		return fmt.Errorf("Failed to create certificate: %s", err)
	}

	return writeToFile(b, key, certFilename, keyFilename)
}

// ==================================
// ====== Interface & Function ======
// ==================================

type PrivateKey interface {
	Public() crypto.PublicKey
}

type PrivateKeyGenerator func() (PrivateKey, error)

// ECDSA 256 is believed to be better than RSA (May 2022)
func ECDSA() (PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func RSA() (PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

// ====================
// ====== Helper ======
// ====================

func RootCert(cert x509.Certificate) x509.Certificate {
	cert.SerialNumber = big.NewInt(1)
	cert.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	cert.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	cert.BasicConstraintsValid = true
	cert.IsCA = true
	cert.MaxPathLen = 2
	return cert
}

func IntermediateCert(cert x509.Certificate) x509.Certificate {
	cert.SerialNumber = big.NewInt(2)
	cert.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	cert.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	cert.BasicConstraintsValid = true
	cert.IsCA = true
	cert.MaxPathLen = 1
	return cert
}

func ServerCert(cert x509.Certificate) x509.Certificate {
	cert.SerialNumber = big.NewInt(3)
	cert.KeyUsage = x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign
	cert.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth}
	cert.BasicConstraintsValid = true
	return cert
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func writeToFile(cert []byte, key PrivateKey, certFilename, keyFilename string) error {
	certOut, err := os.Create(certFilename)
	if err != nil {
		return fmt.Errorf("Failed to open %v for writing: %v", certFilename, err)
	}
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: cert}); err != nil {
		return fmt.Errorf("Failed to write data to %v: %v", certFilename, err)
	}
	if err := certOut.Close(); err != nil {
		return fmt.Errorf("Error closing %v: %v", certFilename, err)
	}

	keyOut, err := os.OpenFile(keyFilename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("Failed to open %v for writing: %v", keyFilename, err)
	}
	privBytes, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		return fmt.Errorf("Unable to marshal private key: %v", err)
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return fmt.Errorf("Failed to write data to %v: %v", keyFilename, err)
	}
	if err := keyOut.Close(); err != nil {
		return fmt.Errorf("Error closing %v: %v", keyFilename, err)
	}
	return nil
}

func loadFromFile(certFilename, keyFilename string) (*x509.Certificate, PrivateKey, error) {
	// Loading certificate
	certBytes, err := os.ReadFile(certFilename)
	if err != nil {
		return nil, nil, err
	}
	certPem, _ := pem.Decode(certBytes)
	cert, err := x509.ParseCertificate(certPem.Bytes)
	if err != nil {
		return nil, nil, err
	}

	// Loading private key
	keyBytes, err := os.ReadFile(keyFilename)
	if err != nil {
		return nil, nil, err
	}
	keyPem, _ := pem.Decode(keyBytes)
	key, err := x509.ParsePKCS8PrivateKey(keyPem.Bytes)
	if err != nil {
		return nil, nil, err
	}
	return cert, key.(PrivateKey), nil
}
