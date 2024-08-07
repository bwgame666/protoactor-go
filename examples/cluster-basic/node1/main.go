package main

import (
	"fmt"

	console "github.com/bwgame666/goconsole"
	"github.com/bwgame666/protoactor-go/actor"
	"github.com/bwgame666/protoactor-go/cluster"
	"github.com/bwgame666/protoactor-go/cluster/clusterproviders/consul"
	"github.com/bwgame666/protoactor-go/cluster/identitylookup/disthash"
	"github.com/bwgame666/protoactor-go/remote"
)

func main() {
	c := startNode()

	fmt.Print("\nBoot other nodes and press Enter\n")
	console.ReadLine()
	pid := c.Get("abc", "hello")
	fmt.Printf("Got pid %v", pid)
	res, _ := c.Request("abc", "hello", &actor.Touch{})
	fmt.Printf("Got response %v", res)

	fmt.Println()
	console.ReadLine()
	c.Shutdown(true)
}

func startNode() *cluster.Cluster {
	system := actor.NewActorSystem()

	provider, _ := consul.New()
	lookup := disthash.New()
	config := remote.Configure("localhost", 0)
	clusterConfig := cluster.Configure("my-cluster", provider, lookup, config)
	c := cluster.New(system, clusterConfig)
	c.StartMember()

	return c
}
