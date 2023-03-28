package main

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	greetv1 "github.com/Hirochon/connect-go-test/server/protocolbuffers/greet/v1"
	"github.com/Hirochon/connect-go-test/server/protocolbuffers/greet/v1/greetv1connect"
	"github.com/bufbuild/connect-go"
)

func TestGreetUnaryHandler(t *testing.T) {
	t.Parallel()
	mux := server()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	cases := []struct {
		scenario string
		name     string
		want     string
	}{
		{
			scenario: "Twitterのユーザー名",
			name:     "heacet43",
			want:     "Hello, heacet43!",
		},
		{
			scenario: "GitHubのユーザー名",
			name:     "Hirochon",
			want:     "Hello, Hirochon!",
		},
		{
			scenario: "今の気持ち",
			name:     "お腹すいた",
			want:     "Hello, お腹すいた!",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.scenario, func(t *testing.T) {
			t.Parallel()
			client := greetv1connect.NewGreetServiceClient(
				server.Client(),
				server.URL,
			)
			res, err := client.GreetUnary(context.Background(), connect.NewRequest(&greetv1.GreetUnaryRequest{
				Name: c.name,
			}))
			if err != nil {
				t.Error(err)
			}
			if res.Msg.GetGreeting() != c.want {
				t.Errorf("greeting got: %s, want: %s", res.Msg.GetGreeting(), c.want)
			}
		})
	}
}

func TestGreetServerStreamHandler(t *testing.T) {
	t.Parallel()
	mux := server()
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)
	cases := []struct {
		scenario string
		name     string
	}{
		{
			scenario: "Twitterのユーザー名",
			name:     "heacet43",
		},
		{
			scenario: "GitHubのユーザー名",
			name:     "Hirochon",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.scenario, func(t *testing.T) {
			t.Parallel()
			client := greetv1connect.NewGreetServiceClient(
				server.Client(),
				server.URL,
			)
			stream, err := client.GreetServerStream(context.Background(), connect.NewRequest(&greetv1.GreetServerStreamRequest{
				Name: c.name,
			}))
			if err != nil {
				t.Error(err)
			}
			i := 0
			for stream.Receive() {
				greeting := stream.Msg().GetGreeting()
				if greeting != fmt.Sprintf("Hello, %s! (%d)", c.name, i) {
					t.Errorf("greeting got: %s, want: %s", greeting, fmt.Sprintf("Hello, %s! (%d)", c.name, i))
				}
				i++
			}
		})
	}
}
