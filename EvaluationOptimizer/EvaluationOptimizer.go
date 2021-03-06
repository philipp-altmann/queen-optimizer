package EvaluationOptimizer

import (
	"sort"
	"fmt"
	"os/exec"
	"os"
	"time"
	"math/rand"
	f "github.com/philipp-altmann/QueenOptimizer/field"
)

const fieldSize = 8
const evaluationCycles = 1000
const populationSize = 100
const selectionFactor = 0.6
const recombinationFactor = 0.4
const mutationFactor = 0.3

func Optimize() (progress [] int) {

	//Init random generation
	var generation [] f.Field
	for i := 0; i < populationSize; i++ {
		generation = append(generation, f.GenerateRandom(fieldSize))
	}
	//Evaluate Fitness
	for j := 0; j < populationSize; j++ {
		generation[j].Evaluate()
	}
	//Sort
	sort.Sort(FitnessSorter(generation))
	//TODO sort while evaluating

	progress = append(progress, generation[0].GetFitness())




	cycle := 0
	gotWorse := 0

	for cycle < evaluationCycles {

		//Select
		generation = generation[:populationSize-populationSize*selectionFactor]

		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
		fmt.Printf("Cycle %d, with best at %d \n", cycle, generation[0].GetFitness())
		//printField(generation[0])

		//fmt.Fprintf(writer, "Cycle %d, with best at %d ", cycle, generation[0].fitness)

		//g.Percent = cycle/evaluationCycles
		//ui.Render(g)

		//printField(generation[0])

		//tm.Flush() // Call it every time at the end of rendering

		if (generation[0].GetFitness() == 0) {
			//We are done
			cycle = evaluationCycles

		}

		mutationCount := 0
		//Mutate
		for p := 0; p < len(generation); p++ {

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			mutate := r.Float32()
			//fmt.Printf("random float: %f\n", random)

			if (mutate > 1-mutationFactor) {
				mutationCount ++
				generation[p].Mutate()
			}

		}

		//fmt.Printf("Mutated %d fields\n", mutationCount)

		//Recombine
		for p := 0; p < len(generation); p++ {

			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			recombine := r.Float32()
			//fmt.Printf("random float: %f\n", random)

			if (recombine > 1-recombinationFactor) {
				combineWith := r.Intn(len(generation))
				newField := generation[p].Recombine(generation[combineWith])
				if (len(generation) < populationSize) {
					generation = append(generation, newField)
				}
			}

		}
		//fmt.Printf(", Generation at size %d \n", len(generation))
		//Fillup
		for len(generation) < populationSize {
			generation = append(generation, f.GenerateRandom(fieldSize))
		}
		//fmt.Printf("Generation at lenth %d",len(generation))


		//Evaluate Fitness
		for j := 0; j < populationSize; j++ {
			generation[j].Evaluate()
		}
		//Sort
		sort.Sort(FitnessSorter(generation))


		cycle++
	}

	//Evaluate Fitness
	for j := 0; j < populationSize; j++ {
		generation[j].Evaluate()
	}
	//Select
	//Sort
	sort.Sort(FitnessSorter(generation))
	var best = generation[0];
	//sort.Sort(FieldSorter(best.queens))
	//Print the best
	best.Print()
	fmt.Printf("%v \n", best)
	fmt.Printf("Got worse %d times\n",gotWorse)

	if(progress[len(progress)-1]<generation[0].GetFitness()){
		gotWorse ++;
	}
	progress = append(progress, generation[0].GetFitness())

	return
}


// Fitness sorter
type FitnessSorter []f.Field

func (f FitnessSorter) Len() int           { return len(f) }
func (f FitnessSorter) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f FitnessSorter) Less(i, j int) bool { return f[i].GetFitness() < f[j].GetFitness() }

