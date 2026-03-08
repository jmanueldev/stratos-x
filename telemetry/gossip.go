package telemetry

import (
	"fmt"
	"math/rand"
	"time"
)

type Telemetry struct {
	NodeID string
	CPU float64
	Memory float64
}

func Gossip(peers []string, data Telemetry) {

	rand.Seed(time.Now().Unix())

	for i:=0;i<3;i++{

		p := peers[rand.Intn(len(peers))]

		fmt.Println("sending telemetry to",p,data)
	}
}
