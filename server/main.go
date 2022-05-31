package main

import (
	"context"
	"log"
	"net"

	"github.com/aniv01/hello_world_grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedHelloWorldServer
}

func main() {
	// create a TCP listener
	lis, err := net.Listen("tcp", ":50001")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()

	proto.RegisterHelloWorldServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		log.Fatal(err)
	}
}

func (s server) Add(c context.Context, r *proto.Request) (*proto.Response, error) {
	resp := proto.Response{}
	resp.R = r.GetA() + r.GetB()
	return &resp, nil
}

func (s server) Multiply(c context.Context, r *proto.Request) (*proto.Response, error) {
	resp := proto.Response{}
	resp.R = r.A * r.A
	return &resp, nil
}
