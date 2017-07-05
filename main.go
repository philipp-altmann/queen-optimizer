package main

import a "github.com/philipp-altmann/QueenOptimizer/ApproximationOptimizer"

const fieldSize = 8
const evaluationCycles = 1000
const populationSize = 200
const selectionFactor = 0.6
const recombinationFactor = 0.4
const mutationFactor = 0.3

func main() {
	//e.Optimize()
	a.Optimize()
}
