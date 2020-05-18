package ds

import (
	"fmt"
	"math/rand"
	"time"

	"../util"
)

type Individual struct {
	Genes     []int   // genes sequence
	Distances float32 //sum of the distances
	Fitness   float32 //fitness of the this species
}

func NewDummyIndividual(num int) *Individual {
	return &Individual{
		Genes: make([]int, num),
	}
}

//NewRandomIndividual random generate the genes
func NewRandomIndividual(num int) *Individual {
	genes := make([]int, num)
	for i := 0; i < num; i++ {
		genes[i] = i
	}

	util.Shuffle(genes)
	return &Individual{
		Genes: genes,
	}
}

//CalcFitness calculate the distance and fitness
func (node *Individual) CalcFitness(distMap [][]float32) {
	totDist := float32(0)
	for i := 0; i < len(node.Genes); i++ {
		curNode := node.Genes[i]
		nextNode := node.Genes[(i+1)%len(node.Genes)]
		totDist += distMap[curNode][nextNode]
	}
	node.Distances = totDist
	node.Fitness = 1.0 / totDist
}

//Clone deeply
func (individual *Individual) Clone() *Individual {
	copiedSpeciesNode := NewDummyIndividual(len(individual.Genes))
	copy(copiedSpeciesNode.Genes, individual.Genes)
	copiedSpeciesNode.Distances = individual.Distances
	copiedSpeciesNode.Fitness = individual.Fitness
	return copiedSpeciesNode
}

func (individual *Individual) PrintRoute() {
	for _, gene := range individual.Genes {
		fmt.Printf("%d->", gene)
	}
	fmt.Println(individual.Genes[0])
	fmt.Printf("node distances:%.3f\n", individual.Distances)
}

func (individual *Individual) DoMutate() *Individual {
	var m, n int
	for {
		rand.Seed(time.Now().UnixNano())
		if len(individual.Genes)-2 < 1 {
			panic("oh no, less than 1")
		}
		m = rand.Intn(len(individual.Genes) - 2)
		n = rand.Intn(len(individual.Genes))

		if m < n {
			break
		}
	}

	for i, j := 0, (n-m+1)>>1; i < j; i++ {
		individual.Genes[m+i], individual.Genes[n-i] = individual.Genes[n-i], individual.Genes[m+i]
	}
	return individual
}

func (individual *Individual) PushMutate() *Individual {
	var m, n int
	for {
		rand.Seed(time.Now().UnixNano())
		m = rand.Intn(len(individual.Genes) >> 1)
		n = rand.Intn(len(individual.Genes))

		if m < n {
			break
		}
	}

	mutatedGenes := make([]int, 0, len(individual.Genes))
	mutatedGenes = append(mutatedGenes, individual.Genes[m:n]...)
	mutatedGenes = append(mutatedGenes, individual.Genes[0:m]...)
	mutatedGenes = append(mutatedGenes, individual.Genes[n:]...)
	copy(individual.Genes, mutatedGenes)
	return individual
}
