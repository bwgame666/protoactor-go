package main

import (
	"fmt"
	"golang.org/x/exp/slog"

	"cluster-grain/shared"

	console "github.com/bwgame666/goconsole"
	"github.com/bwgame666/protoactor-go/actor"
	"github.com/bwgame666/protoactor-go/cluster"
	"github.com/bwgame666/protoactor-go/cluster/clusterproviders/consul"
	"github.com/bwgame666/protoactor-go/cluster/identitylookup/disthash"
	"github.com/bwgame666/protoactor-go/remote"
)

type HelloGrain struct{}

func (h HelloGrain) Init(ctx cluster.GrainContext)           {}
func (h HelloGrain) Terminate(ctx cluster.GrainContext)      {}
func (h HelloGrain) ReceiveDefault(ctx cluster.GrainContext) {}

func (h HelloGrain) SayHello(request *shared.HelloRequest, ctx cluster.GrainContext) (*shared.HelloResponse, error) {
	ctx.Logger().Info("SayHello", slog.String("name", request.Name))
	return &shared.HelloResponse{Message: "Hello " + request.Name}, nil
}

func main() {
	system := actor.NewActorSystem()
	provider, _ := consul.New()
	lookup := disthash.New()
	remoteConfig := remote.Configure("localhost", 0)
	helloKind := shared.NewHelloKind(func() shared.Hello {
		return &HelloGrain{}
	}, 0)
	clusterConfig := cluster.Configure("my-cluster", provider, lookup, remoteConfig,
		cluster.WithKinds(helloKind))

	c := cluster.New(system, clusterConfig)
	c.StartMember()
	fmt.Print("\nBoot other nodes and press Enter\n")
	console.ReadLine()
	c.Shutdown(true)
}
