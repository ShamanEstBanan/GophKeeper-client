package client

import (
	"GophKeeper-client/internal/entity"
	"GophKeeper-client/internal/pkg/encrypt"
	pb "GophKeeper-client/internal/proto"
	"errors"
	"fmt"
	"strings"
)

func Create(name, recordType string, data any) error {
	conn, err := NewConnection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()
	c := pb.NewKeeperServiceClient(conn)

	var encryptedData []byte
	switch data.(type) {
	case entity.RecordBankCard:
		encryptedData, err = encryptBankCardInfo(data.(entity.RecordBankCard))
	case entity.RecordPassword:
		encryptedData, err = encryptPasswordInfo(data.(entity.RecordPassword))
	case entity.RecordByteString:
		encryptedData = encryptByteString(data.(entity.RecordByteString))
	case entity.RecordText:
		encryptedData, err = encryptText(data.(entity.RecordText))
	default:
		return errors.New("error type of data. Sorry")
	}
	if err != nil {
		return fmt.Errorf("internal error")
	}
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
	fmt.Printf("	-Type: %s\n", resp.Type)

	return nil
}

func encryptBankCardInfo(data entity.RecordBankCard) ([]byte, error) {
	cardFields := make([]string, 4, 4)

	cardFields[0] = data.Number
	cardFields[1] = data.UserName
	cardFields[2] = data.ExpiredDate
	cardFields[3] = data.CVV

	cardInfo := strings.Join(cardFields, ";")
	encryptedData, err := encrypt.EncryptString([]byte(cardInfo))
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}

func encryptPasswordInfo(data entity.RecordPassword) ([]byte, error) {
	passwordFields := make([]string, 2, 2)

	passwordFields[0] = data.Login
	passwordFields[1] = data.Password

	passwordInfo := strings.Join(passwordFields, ";")
	encryptedData, err := encrypt.EncryptString([]byte(passwordInfo))
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}

func encryptText(data entity.RecordText) ([]byte, error) {
	encryptedData, err := encrypt.EncryptString([]byte(data.Data))
	if err != nil {
		return nil, err
	}
	return encryptedData, nil
}

// TODO надо бы из бинарника читать
func encryptByteString(data entity.RecordByteString) []byte {
	return []byte(data.Data)
}
