package algorithm

import (
	"math"
	"math/rand"
	"time"

	"../constant"
	ds "../data-structure"
	"../util"
)

type GeneticAlgorithm struct {
	globalBest *ds.Individual
	group      *ds.Group
	distMap    [][]float32
}

func NewGeneticAlgorithm() *GeneticAlgorithm {
	return &GeneticAlgorithm{}
}

func (ga *GeneticAlgorithm) Initialize(populationSize, geneSize int, distMap [][]float32) {
	ga.distMap = distMap
	ga.group = ds.NewGroup(populationSize)
	for i := 0; i < populationSize; i++ {
		ga.group.Add(ds.NewRandomIndividual(geneSize))
	}
	localBest := ga.group.GetCurBestIndividual(ga.distMap)
	//init global best
	ga.globalBest = ds.NewDummyIndividual(geneSize)
	ga.globalBest.Distances = float32(math.MaxFloat32)
	ga.updateGlobalBest(localBest)
}

func (ga *GeneticAlgorithm) Develop(developNum int) (best *ds.Individual) {
	for i := 1; i < developNum; i++ {
		ga.nextGeneration()
	}
	return ga.globalBest
}

func (ga *GeneticAlgorithm) nextGeneration() {
	ga.selection()
	ga.crossover()
	ga.mutation()
	localBest := ga.group.GetCurBestIndividual(ga.distMap)
	ga.updateGlobalBest(localBest)
}

func (ga *GeneticAlgorithm) selection() {
	initNum := 4
	parents := make([]*ds.Individual, 0, ga.group.PopulationSize())
	parents = append(parents, ga.group.GetCurBestIndividual(ga.distMap).Clone())
	parents = append(parents, ga.globalBest.Clone().DoMutate())
	parents = append(parents, ga.globalBest.Clone().PushMutate())
	parents = append(parents, ga.globalBest.Clone())

	roulette := setRoulette(ga.group)
	for i := initNum; i < ga.group.PopulationSize(); i++ {
		rand.Seed(time.Now().UnixNano())
		//maybe they will roll out to the same individual, thus need to copy
		parents = append(parents, ga.group.Population[wheelOut(roulette, rand.Float32())].Clone())
	}
	ga.group.Population = parents
}

func (ga *GeneticAlgorithm) crossover() {
	parentsOrders := make([]int, 0, ga.group.PopulationSize())
	for i := 0; i < ga.group.PopulationSize(); i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Float32() < constant.CROSSOVER_PROBABILITY {
			parentsOrders = append(parentsOrders, i)
		}
	}

	util.Shuffle(parentsOrders)
	for i := 0; i < len(parentsOrders)-1; i += 2 {
		ga.doCrossover(parentsOrders[i], parentsOrders[i+1])
	}
}

func (ga *GeneticAlgorithm) doCrossover(x, y int) {
	child1 := ga.getChild(util.NextIndex, x, y)
	child2 := ga.getChild(util.PrevIndex, x, y)
	ga.group.Population[x] = child1
	ga.group.Population[y] = child2
}

func (ga *GeneticAlgorithm) getChild(findIndex func(a []int, index int) int, x, y int) (child *ds.Individual) {
	geneSize := len(ga.group.Population[x].Genes)
	childGenes := make([]int, 0, geneSize)
	px := ga.group.Population[x].Clone()
	py := ga.group.Population[y].Clone()
	rand.Seed(time.Now().UnixNano())
	c := px.Genes[rand.Intn(geneSize)]
	childGenes = append(childGenes, c)

	for len(px.Genes) > 1 {
		dx := px.Genes[findIndex(px.Genes, util.IndexOf(px.Genes, c))]
		dy := py.Genes[findIndex(py.Genes, util.IndexOf(py.Genes, c))]
		px.Genes = util.DeleteByValue(px.Genes, c)
		py.Genes = util.DeleteByValue(py.Genes, c)
		if ga.distMap[c][dx] < ga.distMap[c][dy] {
			c = dx
		} else {
			c = dy
		}
		childGenes = append(childGenes, c)
	}
	return &ds.Individual{
		Genes: childGenes,
	}
}

func (ga *GeneticAlgorithm) mutation() {
	for i := 0; i < ga.group.PopulationSize(); i++ {
		rand.Seed(time.Now().UnixNano())
		if rand.Float32() < constant.MUTATION_PROBABILITY {
			rand.Seed(time.Now().UnixNano())
			if rand.Float32() > 0.5 {
				ga.group.Population[i] = ga.group.Population[i].DoMutate()
			} else {
				ga.group.Population[i] = ga.group.Population[i].PushMutate()
			}
			//i--
		}
	}
}

func (ga *GeneticAlgorithm) updateGlobalBest(localBest *ds.Individual) {
	if localBest.Distances < ga.globalBest.Distances {
		ga.globalBest = localBest.Clone()
	}
}

func setRoulette(group *ds.Group) (roulette []float32) {
	roulette = make([]float32, group.PopulationSize())

	//set the roulette
	sum := float32(0)
	for _, individual := range group.Population {
		sum += individual.Fitness
	}

	for i, individual := range group.Population {
		roulette[i] = individual.Fitness / sum
	}

	for i := 1; i < len(roulette); i++ {
		roulette[i] += roulette[i-1]
	}
	return
}

func wheelOut(roulette []float32, rand float32) int {
	for i := 0; i < len(roulette); i++ {
		if rand <= roulette[i] {
			return i
		}
	}
	return -1
}
