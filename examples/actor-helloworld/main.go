package main

import (
	console "github.com/bwgame666/goconsole"
	"github.com/bwgame666/protoactor-go/actor"
	"log/slog"
)

type (
	hello      struct{ Who string }
	helloActor struct{}
)

func (state *helloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *hello:
		context.Logger().Info("Hello ", slog.String("who", msg.Who))
	}
}

func main() {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(func() actor.Actor { return &helloActor{} })

	pid := system.Root.Spawn(props)
	system.Root.Send(pid, &hello{Who: "Roger"})
	_, _ = console.ReadLine()
}
