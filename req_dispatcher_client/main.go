package main

import (
	"context"
	"log"
	"strconv"
	"time"

	pb "example.com/grpc-sample/req_dispatcher"
	"google.golang.org/grpc"
)

const (
	host = "localhost"
	port = 8080
	path = "/api/example"
)

func main() {
	conn, err := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewReqDispatcherClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Dispatch(ctx, &pb.ReqDispatchRequest{Host: "example.com", Port: 8080, Path: "/api/example"})
	if err != nil {
		log.Fatalf("could not dispatch: %v", err)
	}
	log.Println("===Result========")
	log.Printf("status: %v", r.GetStatus())
	log.Printf("message: %v", r.GetMessage())
	log.Printf("latencyMs: %v", r.GetLatencyMs())
	log.Println("=================")
}
