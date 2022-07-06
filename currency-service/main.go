package main

import (
	"fmt"
	"net"
	"os"

	hclog "github.com/hashicorp/go-hclog"
	protos "github.com/rabadiyaronak/microservice-go/currency-service/protos/currency"
	"github.com/rabadiyaronak/microservice-go/currency-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	//createe gRpc server
	gs := grpc.NewServer()

	//create instance of currenct server
	cs := server.NewCurrency(log)

	//register currency to gRpc server

	protos.RegisterCurrencyServer(gs, cs)

	//Register the reflection service which allows client to query list of the methods
	//supported by server
	reflection.Register(gs)

	//Create TCP socket for inbound request
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))

	if err != nil {
		log.Error("Unable to create Listner", "error", err)
		os.Exit(-1)
	}

	//listen for request
	gs.Serve(l)
}
