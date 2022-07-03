package main

import (
	"context"
	server "courseselection/kitex_gen/Server"
	"courseselection/kitex_gen/Server/service"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
)

func main() {
	cli, err := service.NewClient("course.selection", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	req := &server.LoginRequest{
		Username: "22070301001",
		Password: "22070301001",
	}
	resp, err := cli.Login(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(resp.Message)
	fmt.Println(*resp.Authority)
}
