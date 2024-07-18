package main

import (
	"fmt"
	"time"

	console "github.com/bwgame666/goconsole"
	"github.com/bwgame666/protoactor-go/actor"
)

type (
	hello      struct{ Who string }
	helloActor struct{}
)

func (state *helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		fmt.Printf("Hello %s\n", msg.Who)
	}
}

func main() {
	system := actor.NewActorSystem(actor.WithDefaultPrometheusProvider(2222))
	props := actor.PropsFromProducer(func() actor.Actor {
		return &helloActor{}
	})

	pid := system.Root.Spawn(props)
	system.Root.Request(pid, &hello{Who: "Prometheus Exporter"})
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Visit http://localhost:2222")
	_, _ = console.ReadLine()
}
