package certhandling

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc/credentials"
)

func LoadTLSServerCredentials(cert, key string) (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	wd, err := os.Getwd()
	log.Infof("Working directory: %s", wd)
	log.Infof("cert: %s", cert)
	log.Info(filepath.Join(wd, cert))
	serverCert, err := tls.LoadX509KeyPair(filepath.Join(wd, cert), filepath.Join(wd, key))
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func LoadTLSClientCredentials(caCert string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	wd, err := os.Getwd()
	log.Infof("Working directory: %s", wd)
	log.Infof("caCert: %s", caCert)
	log.Info(filepath.Join(wd + caCert))
	pemServerCA, err := ioutil.ReadFile(filepath.Join(wd, caCert))
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}
