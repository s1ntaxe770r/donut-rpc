package main

import (
	"context"
	"log"
	"net"

	"github.com/s1ntaxe770r/donut-rpc/db"
	pb "github.com/s1ntaxe770r/donut-rpc/proto"
	"github.com/s1ntaxe770r/donut-rpc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5050"
)

var (
	lg   = utils.NewDonutLogger()
	conn = db.Connect()
)

type DonutServer struct {
	pb.UnimplementedDonutShopServer
	lg *log.Logger
}

func (ds *DonutServer) GetDonut(ctx context.Context, in *pb.DonutRequest) (*pb.Donut, error) {
	donut, err := db.GetDonut(conn, in)
	ds.lg.Println(in)
	if err != nil {
		return nil, err
	}
	return donut, nil
}

func (ds *DonutServer) MakeDonut(ctx context.Context, in *pb.Donut) (*pb.DonutRequest, error) {
	_, err := db.MakeDonut(conn, in)
	if err != nil {
		return nil, err
	}
	ds.lg.Println(in)
	return &pb.DonutRequest{Name: in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		lg.Panicf("could not open shop %s", err.Error())
	}

	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterDonutShopServer(server, &DonutServer{lg: lg})
	lg.Printf("shop opened on %v", lis.Addr())

	err = server.Serve(lis)
	if err != nil {
		lg.Panicf("oops couldn't open up shop %s", err.Error())
	}

}
