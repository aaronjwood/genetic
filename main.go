package main

import (
	"github.com/aaronjwood/genetic/individual"
	"github.com/aaronjwood/genetic/population"
	"fmt"
)

func main() {
	const size int64 = 10000

	pop := population.New(size)
	indvs := individual.New(size)

	fmt.Println(pop.GetFitness())
	fmt.Println(indvs.GetFitness())

}
