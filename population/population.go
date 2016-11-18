package population

import (
	"geneticalgos/individual"
	"math/rand"
)

type Population struct {
	Elite          int
	PopulationSize int
	Rounds         int
	Mutation       float64
	Crossover      float64
	generator      *rand.Rand
	population     []individual.Individual
	fitness        float64
}

//Creates a new population
func New(size int) *Population {
	pop := &Population{
		population: make([]individual.Individual, size),
	}

	for i := range pop.population {
		append(pop.population, individual.New(1000).GenerateGenes())
	}
}
