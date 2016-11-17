package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Individual struct {
	Size      int64
	genes     []int64
	fitness   int64
	generator *rand.Rand
}

//Create a new Individual
func (i *Individual) New(size int64) *Individual {
	return &Individual{
		Size:    size,
		genes:   make([]int64, size),
		fitness: 0,
	}
}

//Gets the current fitness level
func (i *Individual) getFitness() int64 {
	return i.fitness
}

//Sets the fitness level to be used
func (i *Individual) setFitness(fitness int64) *Individual {
	i.fitness = fitness
	return i
}

//Gets a gene at a specified index
func (i *Individual) getGene(idx int64) int64 {
	return i.genes[idx]
}

func (i *Individual) getGenes() []int64 {
	return i.genes
}

//Sets a gene at a specified index
func (i *Individual) setGene(idx, gene int64) *Individual {
	i.genes[idx] = gene

	return i
}

//Generates random genes and fills up the gene pool
func (i *Individual) generateGenes() *Individual {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	for c := range i.genes {
		i.setGene(int64(c), generator.Int63n(2))
	}

	return i
}

//Mutates a gene by flipping it
func (i *Individual) mutate() *Individual {
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	idx := generator.Int63n(i.Size)
	i.setGene(idx, 1-i.getGene(idx))

	return i
}

//Determines a fitness value based on all genes
func (i *Individual) evaluateFitness() int64 {
	i.fitness = 0
	for c := range i.genes {
		i.fitness += i.getGene(int64(c))
	}

	return i.fitness
}

func main() {
	indv := new(Individual).New(10000)
	indv.generateGenes()
	fmt.Println(indv.evaluateFitness())
}
