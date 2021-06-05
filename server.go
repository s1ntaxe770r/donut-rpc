package main

import (
	"context"
	"log"
	"net"

	pb "github.com/s1ntaxe770r/donut-rpc/proto"
	"github.com/s1ntaxe770r/donut-rpc/utils"
	"google.golang.org/grpc"
)

const (
	port = ":5050"
)

var (
	lg = utils.NewDonutLogger()
)

type DonutServer struct {
	pb.UnimplementedDonutShopServer
}

func (ds *DonutServer) GetDonut(ctx context.Context, in *pb.DonutRequest) (*pb.Donut, error) {
	lg.Println(in.GetName())
	return &pb.Donut{Name: "hey"}, nil
}

func (ds *DonutServer) MakeDonut(ctx context.Context, in *pb.Donut) (*pb.DonutRequest, error) {
	lg.Println(in.GetName())
	return &pb.DonutRequest{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not open shop %s", err.Error())
	}
	server := grpc.NewServer()
	pb.RegisterDonutShopServer(server, &DonutServer{})
	lg.Printf("shop opened on %v", lis.Addr())

	err = server.Serve(lis)
	if err != nil {
		lg.Panicf("oops couldn't open up shop %s", err.Error())
	}

}
