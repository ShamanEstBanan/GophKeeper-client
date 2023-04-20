package client

import (
	"GophKeeper-client/internal/config"
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
	"os"
)

func NewConnection() (*grpc.ClientConn, error) {
	cfg := config.GetConfig()
	conn, err := grpc.Dial(cfg.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("create connnection: %w", err)
	}
	return conn, nil
}

func buildContextWithToken() (*context.Context, error) {
	cfg := config.GetConfig()
	tokenFile := cfg.Directory + "/" + cfg.TokenFile
	token, err := os.ReadFile(tokenFile)
	if err != nil {
		fmt.Println("Error reading file with token:", err)
		return nil, err
	}

	md := metadata.Pairs("jwt-token", string(token))
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	return &ctx, nil
}

func saveMasterPass(password string) error {
	encryptKey := sha256.Sum256([]byte(password))
	cfg := config.GetConfig()
	passwordFile := cfg.Directory + "/master_key"
	if err := os.WriteFile(passwordFile, encryptKey[:], 0777); err != nil {
		return err
	}
	return nil
}
func removeMasterPass() error {
	cfg := config.GetConfig()
	passwordFile := cfg.Directory + "/master_key"

	if err := os.Remove(passwordFile); err != nil {
		return err
	}
	return nil
}
func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	cfg := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(cfg), nil
}
