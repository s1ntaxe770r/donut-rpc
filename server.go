package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	lg              = utils.NewDonutLogger()
	conn            = db.Connect()
	rpcMetrics      = grpc_prometheus.NewServerMetrics()
	reg             = prometheus.NewRegistry()
	GetDonutCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "get_donut_request_counter",
		Help: "number of requests served by the get donut method",
	}, []string{"getDonut"})
	MakeDonutCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "make_donut_request_counter",
		Help: "number of requests served by the make donut method",
	}, []string{"MakeDonut"})
	GetDonutsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "get_donust_request_counter",
		Help: "number of requests served by the get donuts method",
	}, []string{"getDonuts"})
)

func init() {
	reg.MustRegister(rpcMetrics, GetDonutCounter, GetDonutsCounter, MakeDonutCounter)

}

type DonutServer struct {
	pb.UnimplementedDonutShopServer
	lg *log.Logger
}

func (ds *DonutServer) GetDonut(ctx context.Context, in *pb.DonutRequest) (*pb.Donut, error) {
	donut, err := db.GetDonut(conn, in)
	GetDonutCounter.WithLabelValues(in.Name).Inc()
	ds.lg.Println(in)
	if err != nil {
		return nil, err
	}
	return donut, nil
}

func (ds *DonutServer) GetDonuts(ctx context.Context, in *pb.DonutRequest) (*pb.Donuts, error) {
	donuts, err := db.GetDonuts(conn)
	GetDonutsCounter.WithLabelValues(in.Name).Inc()
	if err != nil {
		return nil, err
	}
	return donuts, nil

}
func (ds *DonutServer) MakeDonut(ctx context.Context, in *pb.Donut) (*pb.DonutRequest, error) {
	_, err := db.MakeDonut(conn, in)
	MakeDonutCounter.WithLabelValues(in.Name).Inc()
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

	server := grpc.NewServer(
		grpc.StreamInterceptor(rpcMetrics.StreamServerInterceptor()),
		grpc.UnaryInterceptor(rpcMetrics.UnaryServerInterceptor()),
	)
	httpServer := &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%d", 9092)}
	reflection.Register(server)
	pb.RegisterDonutShopServer(server, &DonutServer{lg: lg})
	rpcMetrics.InitializeMetrics(server)
	lg.Printf("shop opened on %v", lis.Addr())

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("Unable to start a http server.")
		}
	}()
	err = server.Serve(lis)
	if err != nil {
		lg.Panicf("oops couldn't open up shop %s", err.Error())
	}

}
