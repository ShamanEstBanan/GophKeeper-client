package client

import (
	"GophKeeper-client/internal/entity"
	pb "GophKeeper-client/internal/proto"
	"errors"
	"fmt"
	"strings"
)

func CreatePassword(name, password, recordType string) error {
	conn, err := NewConnection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	c := pb.NewKeeperServiceClient(conn)

	req := &pb.CreateRecordRequest{
		Name: name,
		Type: recordType,
		Data: []byte(password),
	}

	ctx, err := buildContextWithToken()
	if err != nil {
		return err
	}
	resp, err := c.CreateRecord(*ctx, req)
	if err != nil {
		return err
	}
	fmt.Printf("Created record:\n")
	fmt.Printf("	-ID: %s\n", resp.Id)
	fmt.Printf("	-Name: %s\n", resp.Name)
	fmt.Printf("	-Password: %s\n", resp.Data)
	fmt.Printf("	-Type: %s\n", resp.Type)

	return nil
}

func Create(name, recordType string, data any) error {
	conn, err := NewConnection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	var encryptedData []byte
	switch data.(type) {
	case entity.RecordBankCard:
		encryptedData = encryptBankCardInfo(data.(entity.RecordBankCard))
	case entity.RecordPassword:
		encryptedData = encryptPasswordInfo(data.(entity.RecordPassword))
	case entity.RecordByteString:
		encryptedData = []byte{}
	case entity.RecordText:
		encryptedData = []byte{}

	default:
		return errors.New("error type of data. Sorry")
	}

	c := pb.NewKeeperServiceClient(conn)

	req := &pb.CreateRecordRequest{
		Name: name,
		Type: recordType,
		Data: encryptedData,
	}

	ctx, err := buildContextWithToken()
	if err != nil {
		return err
	}
	resp, err := c.CreateRecord(*ctx, req)
	if err != nil {
		return err
	}
	fmt.Printf("Created record:\n")
	fmt.Printf("	-ID: %s\n", resp.Id)
	fmt.Printf("	-Name: %s\n", resp.Name)
	fmt.Printf("	-Data: %s\n", resp.Data)
	fmt.Printf("	-Type: %s\n", resp.Type)

	return nil
}

func encryptBankCardInfo(data entity.RecordBankCard) []byte {
	cardFields := make([]string, 4, 4)

	cardFields[0] = data.Number
	cardFields[1] = data.UserName
	cardFields[2] = data.ExpiredDate
	cardFields[3] = data.CVV

	cardInfo := strings.Join(cardFields, ";")

	return []byte(cardInfo)
}

func encryptPasswordInfo(data entity.RecordPassword) []byte {
	passwordFields := make([]string, 2, 2)

	passwordFields[0] = data.Login
	passwordFields[1] = data.Password

	passwordInfo := strings.Join(passwordFields, ";")
	return []byte(passwordInfo)
}
