package main

import (
	"fmt"
	"geneticalgos/individual"
)

func main() {
	indv := individual.New(1000)
	indv.GenerateGenes().EvaluateFitness()
	fmt.Println(indv.GetFitness())
}
