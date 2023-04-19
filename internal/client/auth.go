package client

import (
	"GophKeeper-client/internal/config"
	pb "GophKeeper-client/internal/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
)

func SignUp(login, password string) error {
	// Устанавливаем соединение с сервером
	//Соединение с сервером устанавливается при вызове функции grpc.Dial(). В первом параметре указывается адрес сервера, далее перечисляются опциональные параметры.
	//conn, err := NewConnection(c.addr)
	conn, err := NewConnection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	c := pb.NewAuthServiceClient(conn)
	req := &pb.SignUpRequest{
		Login:    login,
		Password: password,
	}
	_, err = c.SignUp(context.Background(), req)
	if err != nil {
		fmt.Println("Error sign up:", err)
		return err
	}

	fmt.Printf("Success sign request with login:%s and pass:%s\n", login, password)
	return err
}

func Login(login, password string) {
	conn, err := NewConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	c := pb.NewAuthServiceClient(conn)
	req := &pb.LogInRequest{
		Login:    login,
		Password: password,
	}
	var header metadata.MD

	_, err = c.LogIn(context.Background(), req, grpc.Header(&header))
	if err != nil {
		return
	}

	token := header.Get("jwt-token")
	if len(token) < 1 {
		fmt.Println("No token in answer Login request")
		return
	}

	err = writeTokenToFile(token[0])
	if err != nil {
		return
	}
	fmt.Printf("Success Login  with login:%s and pass:%s\n", login, password)
}

func writeTokenToFile(token string) error {
	cfg := config.GetConfig()
	tokenFile := cfg.Directory + "/" + cfg.TokenFile
	err := os.WriteFile(tokenFile, []byte(token), 0o600)
	if err != nil {
		fmt.Println("Error saving token to file:", err)
		return err
	}
	return nil

}
