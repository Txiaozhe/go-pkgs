package handler

import (
	"context"
	"errors"
	"go-pkgs/go-kit/stringsvc/model"
	"strings"

	"github.com/go-kit/kit/endpoint"
)

type StringHandler struct{}

func (StringHandler) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("string is empty")
	}

	return strings.ToUpper(s), nil
}

func (StringHandler) Count(s string) int {
	return len(s)
}

func MakeUppercaseEndpoint(svc model.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return model.UppercaseResponse{
				V:   v,
				Err: err.Error(),
			}, nil
		}

		return model.UppercaseResponse{
			V:   v,
			Err: "",
		}, nil
	}
}

func MakeCountEndpoint(svc model.StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(model.CountRequest)
		v := svc.Count(req.S)
		return model.CountResponse{
			V: v,
		}, nil
	}
}
