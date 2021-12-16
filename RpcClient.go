package main

import (
	"context"
	"fmt"
	"go_project16/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		fmt.Println("rpc连接服务端异常", err.Error())
		return
	}
	defer conn.Close()

	client := proto.NewPersonServiceClient(conn)
	personResponse, err := client.GetPerson(context.Background(), &proto.PersonRequest{
		Id: "1",
	})

	fmt.Println("personResponse===", personResponse, err)

	listPersonResponse, err := client.ListPerson(context.Background(), &proto.PersonRequest{})
	fmt.Println("listPersonResponse===", listPersonResponse, err)

}
