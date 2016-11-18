package main

import (
	"geneticalgos/individual"
	"fmt"
)

func main() {
	indv := individual.New(1000)
	indv.GenerateGenes().EvaluateFitness()
	fmt.Println(indv.GetFitness())
}

