package main

import (
	"fmt"

	"distributedchannels/messages"

	console "github.com/bwgame666/goconsole"
	"github.com/bwgame666/protoactor-go/actor"
	"github.com/bwgame666/protoactor-go/remote"
)

func newMyMessageSenderChannel(context actor.SenderContext) chan<- *messages.MyMessage {
	channel := make(chan *messages.MyMessage)
	remoteChannel := actor.NewPID("127.0.0.1:8080", "MyMessage")
	go func() {
		for msg := range channel {
			context.Send(remoteChannel, msg)
		}
	}()

	return channel
}

func main() {
	system := actor.NewActorSystem()
	remoteConfig := remote.Configure("127.0.0.1", 0)
	remoting := remote.NewRemote(system, remoteConfig)
	remoting.Start()

	channel := newMyMessageSenderChannel(system.Root)

	for i := 0; i < 10; i++ {
		message := &messages.MyMessage{
			Message: fmt.Sprintf("hello %v", i),
		}
		channel <- message
	}

	console.ReadLine()
}
