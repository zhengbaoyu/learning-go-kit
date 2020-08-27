package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"go-kit/service"
)

type UserRequest struct {
	Uid int `json:"uid"`
}

type UserResponse struct {
	Name string `json:"name"`
}

func GenUserEndponit(userService service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		respName := userService.GetUserName(r.Uid)
		return UserResponse{Name: respName}, nil
	}
}
