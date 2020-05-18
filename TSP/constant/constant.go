package constant

type InitStrategy string

const (
	EPS = 1e-6

	POPULATION_SIZE = 30
	DEVELOP_NUM     = 1000

	CROSSOVER_PROBABILITY = 0.9 //probability of crossover
	MUTATION_PROBABILITY  = 0.01
)

var (
	GENE_NUM int
)
