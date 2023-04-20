package client

import (
	pb "GophKeeper-client/internal/proto"
	"fmt"
)

func GetAllRecords() error {
	conn, err := NewConnection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	c := pb.NewKeeperServiceClient(conn)
	req := &pb.GetAllRecordsRequest{}

	ctx, err := buildContextWithToken()
	if err != nil {
		return err
	}
	resp, err := c.GetAllRecords(*ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Your records:")
	for k, v := range resp.Records {
		fmt.Printf("%d. ", k+1)
		fmt.Printf(`ID: %s, Name: "%s", Type:"%s"`, v.Id, v.Name, v.Type)
		fmt.Println()
	}
	return nil
}

func GetAllRecordsByType(recordType string) error {
	conn, err := NewConnection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()

	c := pb.NewKeeperServiceClient(conn)
	req := &pb.GetRecordsByTypeRequest{
		Type: recordType,
	}

	ctx, err := buildContextWithToken()
	if err != nil {
		return err
	}
	resp, err := c.GetRecordsByType(*ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Your records:")
	for k, v := range resp.Records {
		fmt.Printf("%d. ", k+1)
		fmt.Printf(`ID: %s, Name: "%s", Type:"%s"`, v.Id, v.Name, v.Type)
		fmt.Println()
	}
	return nil
}
