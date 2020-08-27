package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	Uid int `json:"uid"`
}

type UserResponse struct {
	Name string `json:"name"`
}

func GenUserEndponit() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return UserResponse{Name: ""}, nil
	}
}
