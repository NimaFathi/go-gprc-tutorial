package main

import (
	"context"
	//"context"
	"flag"
	"fmt"
	"log"
	"time"
	_ "time"

	"github.com/NimaFathi/go-grpc-tutorial/pkg/appdetail"
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

	c := appdetail.NewAppDetailClient(conn)
	fmt.Print(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req1 := appdetail.GetAppDetailRequest{PackageName: "hello"}
	res1, err := c.GetAppDetail(ctx, &req1)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	log.Printf("Create result: <%+v>\n\n", res1)

}
