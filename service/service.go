package main

import (
	"flag"

	"github.com/rpcx-ecosystem/rpcx-demo/service/product"
	"github.com/smallnest/rpcx/server"
)

var (
	addr = flag.String("addr", "localhost:8972", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	s.RegisterName("ProductImage", product.New("./product/static"), "")
	err := s.Serve("tcp", *addr)
	if err != nil {
		panic(err)
	}
}
