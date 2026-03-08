package scheduler

import "fmt"

func ScheduleGPUAware(nodes []Node, mission Mission) []Node {
	selected := []Node{}
	for _, n := range nodes {
		if n.GPUCount >= mission.GPURequired && n.RDMAEnabled {
			selected = append(selected, n)
		}
		if len(selected) >= mission.NodesNeeded {
			break
		}
	}
	fmt.Printf("GPU/RDMA-aware scheduling done for %s\n", mission.Name)
	return selected
}
