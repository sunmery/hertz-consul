package service

import (
	"context"
	v1 "hertz-consul/kitex_gen/hello/v1"
)

type HelloService struct {
	ctx context.Context
} // NewHelloService new HelloService
func NewHelloService(ctx context.Context) *HelloService {
	return &HelloService{ctx: ctx}
}

// Run create note info
func (s *HelloService) Run(req *v1.Req) (resp *v1.Res, err error) {
	// Finish your business logic.

	return
}
