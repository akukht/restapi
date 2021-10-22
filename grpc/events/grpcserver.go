package restapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	proto "restapi/grpc/proto"
	"restapi/pkg/model"

	db "restapi/grpc/repository/postgres"

	_ "github.com/lib/pq"
)

//GRPCServer ...
type GRPCServer struct{}

//GetEvent ...
func (s *GRPCServer) GetEvent(ctx context.Context, req *proto.GetEventRequest) (*proto.GetEventResponse, error) {

	event, err := db.GetEventById(req.GetX())

	if err != nil {
		fmt.Println(err)
	}

	response, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
	}

	return &proto.GetEventResponse{Result: string(response)}, nil
}

func (s *GRPCServer) Auth(ctx context.Context, req *proto.AuthRequest) (*proto.AuthResponse, error) {

	token := model.AuthUserDB(req.GetLogin(), req.GetPassword())

	if token == "" {
		return &proto.AuthResponse{Token: ""}, errors.New("wrong login or password")
	}

	return &proto.AuthResponse{Token: token}, nil
}
