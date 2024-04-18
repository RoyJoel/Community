package main

import (
	"github.com/zhangjiacheng-iHealth/IHCommunity/cmd/grpc"
	"github.com/zhangjiacheng-iHealth/IHCommunity/cmd/web"
	"github.com/zhangjiacheng-iHealth/IHCommunity/cmd/aws"
)

func main() {
	go grpc.Run()
	go web.Run()
	go aws.Run()
	select {}
}
