package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	apiVersion = "v1"
)

func main() {
	address := flag.String("server", "", "gRPC server in format host:port")
	flag.Parse()
	conn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewAppDetailClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	t := time.Now().In(time.UTC)
	pfx := t.Format(time.RFC3339Nano)
}
