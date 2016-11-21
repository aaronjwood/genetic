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

	for i := 0; i < size; i++ {
		pop.population = append(pop.population, individual.New(1000).GenerateGenes())
	}

	return pop
}

//Specifies the population of individuals to use
func (p *Population) setPopulation(pop []individual.Individual) *Population {
	copy(p.population, pop)

	return p
}

//Determines the fitness level for the entire population
func (p *Population) evaluate() float64 {
	p.fitness = 0
	for _, v := range p.population {
		p.fitness += v.EvaluateFitness()
	}

	return p.fitness
}
