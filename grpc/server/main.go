package main

import (
	"context"
	"errors"
	"log"
	"net"
	event "restapi/grpc/events"
	api "restapi/grpc/proto"
	"restapi/pkg/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)

	if md["token"] == nil && info.FullMethod != "/api.Event/Auth" {
		return nil, errors.New("you need to authorize")
	}

	//Get token from Headers
	_, err := jwt.GetGWTToken(string(string(md["token"][0])))
	if err != nil {
		return nil, errors.New("token was expired")
	}

	return handler(ctx, req)
}

func streamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Println("---> stream interceptor:", info.FullMethod)
	return handler(srv, stream)
}

func main() {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	srv := &event.GRPCServer{}
	api.RegisterEventServer(s, srv)

	l, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
