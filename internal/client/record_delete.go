package client

import (
	pb "GophKeeper-client/internal/proto"
	"fmt"
)

func Delete(recordID string) error {
	conn, err := NewConnection()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()
	c := pb.NewKeeperServiceClient(conn)

	req := &pb.DeleteRecordRequest{
		Id: recordID,
	}

	ctx, err := buildContextWithToken()
	if err != nil {
		return err
	}
	_, err = c.DeleteRecord(*ctx, req)
	if err != nil {
		return err
	}

	fmt.Printf("Success delete record")

	return nil
}
