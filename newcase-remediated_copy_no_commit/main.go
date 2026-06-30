package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"
	"google.golang.org/grpc"
)

func main() {
	_ = http2.ConfigureServer(&http.Server{}, nil)
	_ = grpc.Dial("")
	fmt.Println("running")
}
