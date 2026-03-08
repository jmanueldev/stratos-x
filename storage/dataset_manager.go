package storage

import "fmt"

type Dataset struct {
	Name string
	Location string
}

func StageDataset(d Dataset, nodes []string){

	for _,n := range nodes {

		fmt.Println("staging dataset",d.Name,"to",n)
	}
}
