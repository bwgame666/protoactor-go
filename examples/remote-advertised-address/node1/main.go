package main

import (
	"log"

	"remoteadvertisedaddress/messages"

	console "github.com/bwgame666/goconsole"
	"github.com/bwgame666/protoactor-go/actor"
	"github.com/bwgame666/protoactor-go/remote"
)

var (
	system      = actor.NewActorSystem()
	rootContext = system.Root
)

func main() {
	cfg := remote.Configure("0.0.0.0", 8081, remote.WithAdvertisedHost("localhost:8081"))
	r := remote.NewRemote(system, cfg)
	r.Start()

	remotePid := actor.NewPID("127.0.0.1:8080", "remote")

	props := actor.
		PropsFromFunc(func(context actor.Context) {
			switch context.Message().(type) {
			case *actor.Started:
				message := &messages.Ping{}
				context.Request(remotePid, message)

			case *messages.Pong:
				log.Println("Received pong from sender")
			}
		})

	rootContext.Spawn(props)

	console.ReadLine()
}
