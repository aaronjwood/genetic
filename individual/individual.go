package individual

import (
	"math/rand"
	"time"
)

type Individual struct {
	genes     []int64
	fitness   int64
	generator *rand.Rand
}

//Create a new Individual
func New(size int64) *Individual {
	return &Individual{
		genes:     make([]int64, size),
		fitness:   0,
		generator: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

//Gets the current fitness level
func (i *Individual) GetFitness() int64 {
	return i.fitness
}

//Sets the fitness level to be used
func (i *Individual) setFitness(fitness int64) *Individual {
	i.fitness = fitness
	return i
}

//Gets a gene at a specified index
func (i *Individual) GetGene(idx int64) int64 {
	return i.genes[idx]
}

func (i *Individual) GetGenes() []int64 {
	return i.genes
}

//Sets a gene at a specified index
func (i *Individual) setGene(idx, gene int64) *Individual {
	i.genes[idx] = gene

	return i
}

//Generates random genes and fills up the gene pool
func (i *Individual) GenerateGenes() *Individual {
	for c := range i.genes {
		i.setGene(int64(c), i.generator.Int63n(2))
	}

	return i
}

//Mutates a gene by flipping it
func (i *Individual) Mutate() *Individual {
	idx := i.generator.Int63n(int64(len(i.genes)))
	i.setGene(idx, 1-i.GetGene(idx))

	return i
}

//Determines a fitness value based on all genes
func (i *Individual) EvaluateFitness() int64 {
	i.fitness = 0
	for c := range i.genes {
		i.fitness += i.GetGene(int64(c))
	}

	return i.fitness
}
