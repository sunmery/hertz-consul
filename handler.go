package main

import (
	"context"
	"github.com/sunmery/hertz-consul/biz/service"
	v1 "github.com/sunmery/hertz-consul/kitex_gen/hello/v1"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// Hello implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) Hello(ctx context.Context, req *v1.Req) (resp *v1.Res, err error) {
	resp, err = service.NewHelloService(ctx).Run(req)

	return resp, err
}
