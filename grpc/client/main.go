package main

import (
	"context"
	"flag"
	"log"
	api "restapi/grpc/proto"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":10000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	c := api.NewEventClient(conn)

	md := metadata.New(map[string]string{"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MzQ3OTg2NzF9.1MJw7tdAn2CSDlKX0sBvck7Ahgm6z-jmKvxKy0nWLBY"})
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Anything linked to this variable will fetch response headers.
	var header metadata.MD

	res, err := c.GetEvent(ctx, &api.GetEventRequest{X: int32(x)}, grpc.Header(&header))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}
