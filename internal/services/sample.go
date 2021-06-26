package services

import (
	"context"
	"haiyon/go-starter/internal/dto"
	"haiyon/go-starter/pkg/xhttp"
)

// Hello .
func (s *Service) Hello(ctx context.Context, body dto.SampleBody) (*xhttp.ResponseException, error) {
	return &xhttp.ResponseException{
		Data: nil,
	}, nil
}
