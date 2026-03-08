package scheduler

import "fmt"

type Node struct {
	ID           string
	GPUCount     int
	GPUMemory    int
	RDMAEnabled  bool
	Zone         string
	Status       string
}

type Mission struct {
	Name        string
	Topology    string
	NodesNeeded int
	GPURequired int
}

func ScheduleMission(nodes []Node, mission Mission) []Node {
	selected := []Node{}
	for _, n := range nodes {
		if n.Status == "idle" && n.GPUCount >= mission.GPURequired && n.RDMAEnabled {
			selected = append(selected, n)
		}
		if len(selected) >= mission.NodesNeeded {
			break
		}
	}
	fmt.Printf("Mission %s scheduled on %d nodes\n", mission.Name, len(selected))
	return selected
}
