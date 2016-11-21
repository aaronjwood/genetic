package population

import (
	"github.com/aaronjwood/genetic/individual"
	"math/rand"
	"time"
)

const (
	Minimization = iota
	Maximization
)

type Population struct {
	Elite      int
	Rounds     int
	Mutation   float64
	Crossover  float64
	generator  *rand.Rand
	population []individual.Individual
	fitness    int64
}

//Creates a new population
func New(size int64) *Population {
	pop := &Population{
		generator: rand.New(rand.NewSource(time.Now().UnixNano())),
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
func (p *Population) rouletteSelection() *individual.Individual {
	selection := p.generator.Int63() * p.fitness
	var idx int
	for idx = 0; idx < len(p.population) && selection > 0; idx++ {
		selection -= p.population[idx].GetFitness()
	}

	return &p.population[idx-1]
}

//Determines the fitness level for the entire population
func (p *Population) evaluate() int64 {
	p.fitness = 0
	for _, v := range p.population {
		p.fitness += v.EvaluateFitness()
	}

	return p.fitness
}

//Finds the best individual based on using minimization or maximization
func (p *Population) BestIndividual(algo int) *individual.Individual {
	var idxMin, idxMax int64 = 0, 0
	var currentMin, currentMax int64 = 1, 0
	var currentVal int64

	for idx := 0; idx < len(p.population); idx++ {
		currentVal = p.population[idx].GetFitness()
		if currentMax < currentMin {
			currentMin = currentVal
			currentMax = currentVal
			idxMin = int64(idx)
			idxMax = int64(idx)
		}

		if currentVal > currentMax {
			currentMax = currentVal
			idxMax = int64(idx)
		}

		if currentVal < currentMin {
			currentMin = currentVal
			idxMin = int64(idx)
		}
	}

	if algo == Minimization {
		return &p.population[idxMin]
	} else if algo == Maximization {
		return &p.population[idxMax]
	}

	return &p.population[idxMax]
}
