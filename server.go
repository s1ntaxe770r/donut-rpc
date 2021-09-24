package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/s1ntaxe770r/donut-rpc/db"
	pb "github.com/s1ntaxe770r/donut-rpc/proto"
	"github.com/s1ntaxe770r/donut-rpc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

const (
	port    = ":5050"
	version = "v0.1"
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
		Name: "get_donuts_request_counter",
		Help: "number of requests served by the get donuts method",
	}, []string{"getDonuts"})
)

func init() {
	reg.MustRegister(rpcMetrics, GetDonutCounter, GetDonutsCounter, MakeDonutCounter)

}

type DonutServer struct {
	pb.UnimplementedDonutShopServer
	lg *log.Logger
	db *gorm.DB
}

func (ds *DonutServer) GetDonut(ctx context.Context, in *pb.DonutRequest) (*pb.Donut, error) {
	donut, err := db.GetDonut(ds.db, in)
	GetDonutCounter.WithLabelValues(in.Name).Inc()
	ds.lg.Println(in)
	if err != nil {
		return nil, err
	}
	defer ds.lg.Printf("%s|handle get donut for %v", time.Now().String(), in)
	return donut, nil
}

func (ds *DonutServer) GetDonuts(ctx context.Context, in *emptypb.Empty) (*pb.Donuts, error) {
	donuts, err := db.GetDonuts(ds.db)
	GetDonutsCounter.WithLabelValues("GetDonuts").Inc()
	if err != nil {
		return nil, err
	}
	defer ds.lg.Printf("%s|handle get donuts", time.Now().String())
	return donuts, nil
}
func (ds *DonutServer) MakeDonut(ctx context.Context, in *pb.Donut) (*pb.DonutRequest, error) {
	_, err := db.MakeDonut(ds.db, in)
	MakeDonutCounter.WithLabelValues(in.Name).Inc()
	if err != nil {
		return nil, err
	}
	defer ds.lg.Printf("%s|handle get make donut for %v", time.Now().String(), in)
	return &pb.DonutRequest{Name: in.GetName()}, nil
}

func (ds *DonutServer) GetVersion(ctx context.Context, in *emptypb.Empty) (*pb.Version, error) {
	return &pb.Version{Number: version}, nil
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
	pb.RegisterDonutShopServer(server, &DonutServer{lg: lg, db: conn})
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
