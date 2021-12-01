package modules

import (
	proto "github.com/meysifk/movies_proto/movies-API"
	"github.com/meysifk/test-backend/second_answer/movies-API/client/config"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

type Module struct {
	Cfg         *config.Config
	MovieClient proto.RouteMovieClient
}

func New(cfg *config.Config) *Module {
	module := Module{}
	module.Cfg = cfg

	module.InitMovieClient()

	return &module
}

func (m *Module) InitMovieClient() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(m.Cfg.Server.MovieServerAddr, opts...)
	if err != nil {
		log.Fatal(err)
	}
	client := proto.NewRouteMovieClient(conn)

	log.Infoln("grpc client for movie initiated")
	m.MovieClient = client

}
