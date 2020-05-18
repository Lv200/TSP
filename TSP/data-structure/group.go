package ds

import (
	"fmt"
	"math"
)

type Group struct {
	Population []*Individual
}

func NewGroup(populationSize int) *Group {
	return &Group{
		Population: make([]*Individual, 0, populationSize),
	}
}

func (group *Group) Add(individual *Individual) {
	if len(group.Population) >= cap(group.Population) {
		panic("group size overflow")
	}

	group.Population = append(group.Population, individual)
}

//GetCurBestIndividual return the best fitness individual in the group
func (group *Group) GetCurBestIndividual(distMap [][]float32) (bestIndividual *Individual) {
	distances := float32(math.MaxFloat32)
	for _, individual := range group.Population {
		individual.CalcFitness(distMap)
		if individual.Distances < distances {
			distances = individual.Distances
			bestIndividual = individual
		}
	}
	return
}

func (group *Group) Traverse() {
	for _, individual := range group.Population {
		individual.PrintRoute()
	}
	fmt.Println("**************")
}

func (group *Group) PopulationSize() int {
	return len(group.Population)
}
