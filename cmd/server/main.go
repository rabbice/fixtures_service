package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/rabbice/fixturesbook/service"
	"google.golang.org/grpc"
	pb "github.com/rabbice/fixturesbook/pb"
)

func main() {
	port := flag.Int("port", 0, "server port")
	flag.Parse()
	log.Print("Started gRPC server on port: ", *port)

	fixtureServer := service.NewFixtureServer(service.NewInMemoryFixtureStore())
	grpcServer := grpc.NewServer()
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
