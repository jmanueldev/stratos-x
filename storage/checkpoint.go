package storage

import (
	"fmt"
	"time"
)

type Checkpoint struct {
	MissionID string
	NodeID    string
	Timestamp time.Time
	Data      string
}

func SaveCheckpoint(missionID, nodeID, data string) Checkpoint {
	cp := Checkpoint{MissionID: missionID, NodeID: nodeID, Timestamp: time.Now(), Data: data}
	fmt.Printf("Checkpoint saved for %s at node %s\n", missionID, nodeID)
	return cp
}

func RecoverCheckpoint(cp Checkpoint) {
	fmt.Printf("Recovered checkpoint for %s at node %s\n", cp.MissionID, cp.NodeID)
}
