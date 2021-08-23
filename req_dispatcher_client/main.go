package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	pb "example.com/grpc-sample/req_dispatcher"
	"google.golang.org/grpc"
)

const (
	defaultHost       = "example.com"
	defaultPort int32 = 443
	defaultPath       = "/api/example"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewReqDispatcherClient(conn)

	host := defaultHost
	if len(os.Args) > 1 {
		host = os.Args[1]
	}
	port := defaultPort
	if len(os.Args) > 2 {
		i, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("could not parse port: %v", err)
		}
		port = int32(i)
	}
	path := defaultPath
	if len(os.Args) > 3 {
		path = os.Args[3]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Dispatch(ctx, &pb.ReqDispatchRequest{Host: host, Port: port, Path: path})
	if err != nil {
		log.Fatalf("could not dispatch: %v", err)
	}
	log.Println("===Result========")
	log.Printf("status: %v", r.GetStatus())
	log.Printf("message: %v", r.GetMessage())
	log.Printf("latencyMs: %v", r.GetLatencyMs())
	log.Println("=================")
}
