package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "example.com/grpc-sample/req_dispatcher"
	"google.golang.org/grpc"
)

const port = ":8080"

type server struct {
	pb.UnimplementedReqDispatcherServer
}

func (s *server) Dispatch(ctx context.Context, in *pb.ReqDispatchRequest) (*pb.ReqDispatchReply, error) {
	log.Printf("Received: %v", in.GetHost()+":"+fmt.Sprintf("%d", in.GetPort())+in.GetPath())
	// TODO: request against specified URL

	var status int32 = 200
	message := "ok"
	var latencyMs int32 = 100

	return &pb.ReqDispatchReply{Status: status, Message: message, LatencyMs: latencyMs}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterReqDispatcherServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
