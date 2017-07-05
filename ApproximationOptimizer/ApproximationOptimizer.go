package ApproximationOptimizer

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sort"
	"time"

	f "github.com/philipp-altmann/QueenOptimizer/field"
	q "github.com/philipp-altmann/QueenOptimizer/queen"
)

const fieldSize = 8
const worstFitness = 120 //todo calc
const evaluationCycles = 0000
const populationSize = 100
const selectionFactor = 0.6
const recombinationFactor = 0.4
const mutationFactor = 0.3
const evaluationPoolSize = 1

func Optimize() (progress []int) {
	var evaluationPool []f.Field
	/*for i := 0; i < evaluationPoolSize; i++ {
		evaluationPool = append(evaluationPool, f.GenerateRandom(fieldSize))
	}*/
	var queens []q.Queen
	queens = append(queens, q.Generate(0, 0))
	queens = append(queens, q.Generate(2, 1))
	queens = append(queens, q.Generate(4, 2))
	queens = append(queens, q.Generate(6, 3))
	queens = append(queens, q.Generate(1, 4))
	queens = append(queens, q.Generate(3, 5))
	queens = append(queens, q.Generate(5, 6))
	queens = append(queens, q.Generate(7, 7))
	evaluationPool = append(evaluationPool, f.Generate(queens, 8))

	var queens2 []q.Queen
	queens2 = append(queens2, q.Generate(0, 0))
	queens2 = append(queens2, q.Generate(0, 0))
	queens2 = append(queens2, q.Generate(0, 0))
	queens2 = append(queens2, q.Generate(0, 0))
	queens2 = append(queens2, q.Generate(0, 0))
	queens2 = append(queens2, q.Generate(0, 0))
	queens2 = append(queens2, q.Generate(0, 0))
	queens2 = append(queens2, q.Generate(0, 0))
	evaluationPool = append(evaluationPool, f.Generate(queens2, 8))

	//Evaluate Pool Fitness
	for j := 0; j < evaluationPoolSize; j++ {
		evaluationPool[j].Evaluate()
		//fmt.Printf("Pool Fitness: %d", evaluationPool[j].GetFitness())
	}
	//Init random generation
	var generation []f.Field
	for i := 0; i < populationSize; i++ {
		generation = append(generation, f.GenerateRandom(fieldSize))
	}
	//Evaluate Fitness
	for j := 0; j < populationSize; j++ {
		generation[j].Approximate(evaluationPool)
		fmt.Print(generation[j].ToPrintable())
		fmt.Printf("Approximated Fitness: %d", generation[j].GetFitness())

		generation[j].Evaluate()
		fmt.Printf("Evaluated Fitness: %d", generation[j].GetFitness())

	}

	// test
	evaluationPool[0].Approximate(evaluationPool)
	fmt.Print(evaluationPool[0].ToPrintable())
	fmt.Printf("\n\nApproximated Fitness: %d", evaluationPool[0].GetFitness())

	//Sort
	sort.Sort(FitnessSorter(generation))
	//TODO sort while evaluating

	progress = append(progress, generation[0].GetFitness())

	cycle := 0
	gotWorse := 0

	for cycle < evaluationCycles {

		//Select
		generation = generation[:populationSize-populationSize*selectionFactor]

		//Mutate
		//mutationCount := 0
		for p := 0; p < len(generation); p++ {

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			mutate := r.Float32()
			if mutate > 1-mutationFactor {
				//mutationCount++
				generation[p].Mutate()
			}
		}
		//fmt.Printf("Mutated %d fields\n", mutationCount)

		//Recombine
		/*for p := 0; p < len(generation); p++ {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			recombine := r.Float32()
			if recombine > 1-recombinationFactor {
				combineWith := r.Intn(len(generation))
				newField := generation[p].Recombine(generation[combineWith])
				if len(generation) < populationSize {
					generation = append(generation, newField)
				}
			}
		}*/
		//fmt.Printf(", Generation at size %d \n", len(generation))

		//Fillup
		for len(generation) < populationSize {
			generation = append(generation, f.GenerateRandom(fieldSize))
		}
		//fmt.Printf("Generation at lenth %d",len(generation))

		//Evaluate Fitness
		for j := 0; j < populationSize; j++ {
			generation[j].Approximate(evaluationPool)
		}
		//Sort
		sort.Sort(FitnessSorter(generation))

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		/*fmt.Printf("Evaluating against\n")
		fmt.Print(evaluationPool[0].ToPrintable())*/
		fmt.Printf("Approximated at: %d\n", generation[0].GetFitness())

		generation[0].Evaluate()
		fmt.Printf("Cycle %d, with best at %d \n", cycle, generation[0].GetFitness())
		fmt.Print(generation[0].ToPrintable())

		cycle++
	}
	time.Sleep(2 * time.Second)

	//Evaluate Fitness
	for j := 0; j < populationSize; j++ {
		generation[j].Evaluate()
	}
	//Select
	//Sort
	sort.Sort(FitnessSorter(generation))
	var best = generation[0]
	//sort.Sort(FieldSorter(best.queens))
	//Print the best
	best.ToPrintable()
	best.Evaluate()

	fmt.Printf("Fitenss at %d \n", best.GetFitness())
	fmt.Printf("%v \n", best)
	fmt.Printf("Got worse %d times\n", gotWorse)

	progress = append(progress, generation[0].GetFitness())

	return

}

// Fitness sorter
type FitnessSorter []f.Field

func (f FitnessSorter) Len() int           { return len(f) }
func (f FitnessSorter) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f FitnessSorter) Less(i, j int) bool { return f[i].GetFitness() < f[j].GetFitness() }
