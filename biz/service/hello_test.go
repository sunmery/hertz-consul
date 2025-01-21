package service

import (
	"context"
	v1 "github.com/sunmery/hertz-consul/kitex_gen/hello/v1"
	"testing"
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
