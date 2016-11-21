package population

import (
	"github.com/aaronjwood/genetic/individual"
	"math/rand"
	"time"
)

type Population struct {
	Elite          int
	Rounds         int
	Mutation       float64
	Crossover      float64
	generator      *rand.Rand
	population     []individual.Individual
	fitness        int64
}

//Creates a new population
func New(size int64) *Population {
	pop := &Population{
		population: make([]individual.Individual, size),
		generator:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}

	for i := int64(0); i < size; i++ {
		pop.population = append(pop.population, *individual.New(size).GenerateGenes())
	}

	pop.evaluate()

	return pop
}

func (p *Population) GetFitness() int64 {
	return p.fitness
}

//Specifies the population of individuals to use
func (p *Population) setPopulation(pop []individual.Individual) *Population {
	copy(p.population, pop)

	return p
}

//Gets the population
func (p *Population) getPopulation() []individual.Individual {
	return p.population
}

//Selects an individual at random
func (p *Population) rouletteSelection() individual.Individual {
	selection := p.generator.Int63() * p.fitness
	var idx int
	for idx = 0; idx < len(p.population) && selection > 0; idx++ {
		selection -= p.population[idx].GetFitness()
	}

	return p.population[idx-1]
}

//Determines the fitness level for the entire population
func (p *Population) evaluate() int64 {
	p.fitness = 0
	for _, v := range p.population {
		p.fitness += v.EvaluateFitness()
	}

	return p.fitness
}
