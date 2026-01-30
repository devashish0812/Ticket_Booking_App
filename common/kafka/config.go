package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"strings"
)

type Config struct {
	Brokers []string
	TLS     *tls.Config
}

func LoadConfig() *Config {
	broker := os.Getenv("KAFKA_BROKER_URL")
	if broker == "" {
		broker = "localhost:9092"
	}

	var tlsConfig *tls.Config

	caString := os.Getenv("KAFKA_CA_PEM")
	certString := os.Getenv("KAFKA_SERVICE_CERT")
	keyString := os.Getenv("KAFKA_SERVICE_KEY")

	if caString != "" && certString != "" && keyString != "" {
		caString = strings.ReplaceAll(caString, `\n`, "\n")
		certString = strings.ReplaceAll(certString, `\n`, "\n")
		keyString = strings.ReplaceAll(keyString, `\n`, "\n")

		// log.Printf("DEBUG CERT START: %q", certString[:50])
		keyPair, err := tls.X509KeyPair([]byte(certString), []byte(keyString))
		if err != nil {
			log.Printf("Failed to load Kafka KeyPair: %v", err)
		} else {
			caCertPool := x509.NewCertPool()
			if ok := caCertPool.AppendCertsFromPEM([]byte(caString)); !ok {
				log.Println("Failed to append Kafka CA Cert")
			} else {
				tlsConfig = &tls.Config{
					Certificates: []tls.Certificate{keyPair},
					RootCAs:      caCertPool,
					MinVersion:   tls.VersionTLS12,
				}
			}
		}
	}

	return &Config{
		Brokers: []string{broker},
		TLS:     tlsConfig,
	}
}
