package client

import (
	"GophKeeper-client/internal/entity"
	"GophKeeper-client/internal/pkg/encrypt"
	pb "GophKeeper-client/internal/proto"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Get(recordID string) error {
	conn, err := NewConnection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()
	c := pb.NewKeeperServiceClient(conn)

	req := &pb.GetRecordRequest{
		Id: recordID,
	}

	ctx, err := buildContextWithToken()
	if err != nil {
		return err
	}
	resp, err := c.GetRecord(*ctx, req)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound, codes.DeadlineExceeded:
				return fmt.Errorf(e.Message())
			default:
				fmt.Println(e.Code(), e.Message())

			}
			fmt.Printf("Не получилось распарсить ошибку %v", err)
			return err
		}
	}
	var originData []byte

	//TODO сделать маппер строки в структуру нужного типа
	switch resp.Type {
	case entity.RecordTypePassword:
		originData, err = decryptText(resp.Data)
	case entity.RecordTypeBankCard:
		originData, err = decryptText(resp.Data)
	case entity.RecordTypeTextString:
		originData, err = decryptText(resp.Data)
	case entity.RecordTypeByteString:
		originData, err = decryptText(resp.Data)

	}
	fmt.Printf("Created record:\n")
	fmt.Printf("	-ID: %s\n", resp.Id)
	fmt.Printf("	-Name: %s\n", resp.Name)
	fmt.Printf("	-Data: %s\n", originData)
	fmt.Printf("	-Type: %s\n", resp.Type)

	return nil
}
func decryptText(data []byte) ([]byte, error) {
	return encrypt.DecryptString(data)
}
