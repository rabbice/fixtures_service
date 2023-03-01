package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/rabbice/fixturesbook/pb"
	"github.com/rabbice/fixturesbook/service"
	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/zpages"
	"google.golang.org/grpc"
)

func main() {

	go func() {
		mux := http.NewServeMux()
		zpages.Handle(mux, "/debug")
		log.Fatal(http.ListenAndServe("localhost:8081", mux))
	}()

	view.RegisterExporter(&exporter.PrintExporter{})

	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatal(err)
	}

	port := flag.Int("port", 0, "server port")
	flag.Parse()
	log.Print("Started gRPC server on port: ", *port)

	fixtureServer := service.NewFixtureServer(service.NewInMemoryFixtureStore())
	grpcServer := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ClientHandler{}))
	pb.RegisterFixtureServiceServer(grpcServer, fixtureServer)

	addr := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("cannot start server:%v ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start server:%v ", err)
	}
}
