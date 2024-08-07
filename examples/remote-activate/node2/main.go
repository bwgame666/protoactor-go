package main

import (
	"runtime"

	"remoteactivate/messages"

	console "github.com/bwgame666/goconsole"
	"github.com/bwgame666/protoactor-go/actor"
	"github.com/bwgame666/protoactor-go/remote"
)

type helloActor struct{}

func (*helloActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *messages.HelloRequest:
		ctx.Respond(&messages.HelloResponse{
			Message: "Hello from remote node",
		})
	}
}

func newHelloActor() actor.Actor {
	return &helloActor{}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	system := actor.NewActorSystem()
	remoteConfig := remote.Configure("127.0.0.1", 8080,
		remote.WithKinds(remote.NewKind("hello", actor.PropsFromProducer(newHelloActor))))

	remoter := remote.NewRemote(system, remoteConfig)
	remoter.Start()

	console.ReadLine()
}
