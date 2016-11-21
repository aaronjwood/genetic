package main

import (
	"fmt"
	"github.com/aaronjwood/genetic/population"
)

func main() {
	const size int64 = 10000

	pop := population.New(size)

	fmt.Println(pop.GetFitness())
	fmt.Println(pop.BestIndividual(population.Maximization).EvaluateFitness())

}
