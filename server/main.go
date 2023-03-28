package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/Hirochon/connect-go-test/server/protocolbuffers/greet/v1/greetv1connect"
)

func server() http.Handler {
	mux := http.NewServeMux()
	mux.Handle(greetv1connect.NewGreetServiceHandler(nil))
	return mux
}

func main() {
	mux := server()
	http.ListenAndServe(
		":5050",
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
