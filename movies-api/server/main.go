package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// logging : for now just use default system logging, will set letter

	// get env form parameter
	var env string
	flag.StringVar(&env, "env", "dev", "set env here eg : dev,stg,prd")
	flag.Parse()

	// init config
	cfg, err := config.InitConfig(env)
	if err != nil {
		log.Fatal("error init config: ", err.Error())
	}

	log.Printf("given config: %+v", cfg)

	flag.Parse()

	// init module
	m := modules.New(cfg)

	// listen tcp
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.Server.Port))
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.ServerOption

	// init grpc
	grpcServer := grpc.NewServer(opts...)
	proto.RegisterRouteMovieServer(grpcServer, m)

	// serve
	log.Println("listen server to port " + fmt.Sprintf("localhost:%d", cfg.Server.Port))
	grpcServer.Serve(lis)
}
