package service

import (
	"context"
	"testing"
	v1 "hertz-consul/kitex_gen/hello/v1"
)

func TestHello_Run(t *testing.T) {
	ctx := context.Background()
	s := NewHelloService(ctx)
	// init req and assert value

	req := &v1.Req{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
