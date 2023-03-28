package main

import (
	"context"
	"fmt"

	greetv1 "github.com/Hirochon/connect-go-test/server/protocolbuffers/greet/v1"
	"github.com/Hirochon/connect-go-test/server/protocolbuffers/greet/v1/greetv1connect"
	"github.com/bufbuild/connect-go"
)

type GreetServer struct {
	greetv1connect.UnimplementedGreetServiceHandler
}

func (s *GreetServer) GreetUnary(
	ctx context.Context,
	req *connect.Request[greetv1.GreetUnaryRequest],
) (*connect.Response[greetv1.GreetUnaryResponse], error) {
	res := connect.NewResponse(&greetv1.GreetUnaryResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	return res, nil
}

func (s *GreetServer) GreetServerStream(
	ctx context.Context,
	req *connect.Request[greetv1.GreetServerStreamRequest],
	stream *connect.ServerStream[greetv1.GreetServerStreamResponse],
) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(&greetv1.GreetServerStreamResponse{
			Greeting: fmt.Sprintf("Hello, %s! (%d)", req.Msg.Name, i),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (s *GreetServer) GreetClientStream(
	ctx context.Context,
	req *connect.ClientStream[greetv1.GreetClientStreamRequest],
) (*connect.Response[greetv1.GreetClientStreamResponse], error) {
	var names string
	for req.Receive() {
		if len(names) > 0 {
			names += ", "
		}
		names += req.Msg().Name
	}
	res := connect.NewResponse(&greetv1.GreetClientStreamResponse{
		Greeting: fmt.Sprintf("Hello, %s!", names),
	})
	return res, nil
}
