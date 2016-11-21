package main

import (
	"fmt"
	"github.com/aaronjwood/genetic/individual"
)

func main() {
	indv := individual.New(1000)
	indv.GenerateGenes().EvaluateFitness()
	fmt.Println(indv.GetFitness())
}
