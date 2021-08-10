package service

import (
	"context"
	"go-starter/common/xhttp"
	"go-starter/internal/dto"
)

// Hello .
func (svc *Service) Hello(ctx context.Context, body dto.SampleBody) (*xhttp.ResponseException, error) {
	return &xhttp.ResponseException{
		Data: nil,
	}, nil
}
