package scheduler

type Topology struct {
	Name        string
	Connections map[string][]string
}

func MeshTopology(nodes []Node) Topology {
	conn := make(map[string][]string)
	for i := range nodes {
		if i+1 < len(nodes) {
			conn[nodes[i].ID] = append(conn[nodes[i].ID], nodes[i+1].ID)
		}
		if i-1 >= 0 {
			conn[nodes[i].ID] = append(conn[nodes[i].ID], nodes[i-1].ID)
		}
	}
	return Topology{Name: "mesh", Connections: conn}
}
