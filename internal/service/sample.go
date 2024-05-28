package service

import (
	"context"
	"go-starter/internal/schema/structs"
	"go-starter/pkg/resp"
	"go-starter/pkg/validator"
)

// Hello .
func (svc *Service) Hello(ctx context.Context, body structs.Sample) (*resp.Exception, error) {
	row, err := svc.sample.Hello(ctx, body)
	if validator.IsNotNil(err) {
		return &resp.Exception{Message: err.Error()}, nil
	}
	return &resp.Exception{
		Data: row,
	}, nil
}
