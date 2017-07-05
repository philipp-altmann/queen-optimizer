package field

import (
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"

	q "github.com/philipp-altmann/QueenOptimizer/queen"
)

type Field struct {
	queens  []q.Queen
	size    int
	fitness int
}

func Generate(queens []q.Queen, size int) (field Field) {
	return Field{queens: queens, size: size, fitness: -1}
}

func GenerateRandom(fieldSize int) (field Field) {
	var queens []q.Queen
	for i := 0; i < fieldSize; i++ {
		queens = append(queens, q.GenerateRandom(fieldSize))
	}
	return Generate(queens, fieldSize)
}

func (field *Field) ToPrintable() (printable string) {
	var fieldArray = make([][]bool, field.size) //[field.size][field.size] bool
	for i := range fieldArray {
		fieldArray[i] = make([]bool, field.size)
	}
	for i := 0; i < field.size; i++ {
		fieldArray[field.queens[i].GetY()][field.queens[i].GetX()] = true
	}
	printable += ("\n+" + strings.Repeat("–––+", len(fieldArray)) + "\n")

	for i := 0; i < len(fieldArray); i++ {
		printable += ("|")
		for j := 0; j < len(fieldArray[i]); j++ {
			if fieldArray[i][j] {
				printable += (" X |")
			} else {
				printable += ("   |")
			}
		}
		printable += ("\n+" + strings.Repeat("–––+", len(fieldArray[i])) + "\n")
	}
	return
}

func (field *Field) Evaluate() {
	fitness := 0
	var worstFitness = field.size * field.size //todo calc

EvaluationLoop:
	for i := 0; i < field.size; i++ {
		for j := i + 1; j < field.size; j++ {
			if field.queens[i].Captures(field.queens[j]) {
				fitness++
				if field.queens[i].Equals(field.queens[j]) {
					fitness = worstFitness
					break EvaluationLoop
				}
			}
		}
	}
	field.fitness = fitness
}

func (field *Field) Approximate(approximationPool []Field) {
	var similarity []float64
	for i := 0; i < len(approximationPool); i++ {
		//similarity = append(similarity, 1-field.Compare(approximationPool[i]))
		similarity = append(similarity, field.Compare(approximationPool[i]))
	}
	var similaritySum float64 = 0
	for i := range similarity {
		similaritySum += similarity[i]
	}

	var approximatefitness float64 = 0
	for i := 0; i < len(approximationPool); i++ {
		approximatefitness += (similarity[i] / similaritySum) * float64(approximationPool[i].fitness)
	}

	field.fitness = int(similaritySum) //int(math.Floor(approximatefitness))

}

func (field *Field) Distance(to Field) (distance int) {
	sort.Sort(FieldSorter(field.queens))
	sort.Sort(FieldSorter(to.queens))

	for i := 0; i < field.size; i++ {
		if !field.queens[i].Equals(to.queens[i]) {
			distance += int(math.Abs(float64(field.queens[i].GetX() - to.queens[i].GetX())))
			distance += int(math.Abs(float64(field.queens[i].GetY() - to.queens[i].GetY())))
		}
	}
	return
}

func (field *Field) Compare(to Field) (difference float64) {
	difference = float64(field.Distance(to))
	difference /= float64(field.size * 2)
	return
}

func (field *Field) GetFitness() (fitness int) {
	return field.fitness
}

func (field1 *Field) Recombine(field2 Field) (field Field) {
	sort.Sort(FieldSorter(field1.queens))
	sort.Sort(FieldSorter(field2.queens))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	cut := 1 + r.Intn(field1.size-2)
	var queens []q.Queen
	for i := 0; i < cut; i++ {
		queens = append(queens, q.Generate(field1.queens[i].GetX(), field1.queens[i].GetY()))
	}
	for j := cut; j < field1.size; j++ {
		queens = append(queens, q.Generate(field2.queens[j].GetX(), field2.queens[j].GetY()))
	}
	return Field{queens: queens, size: field1.size}
}

func (field *Field) Mutate() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var queen = r.Intn(field.size)

	field.queens[queen].Mutate(field.size)
}

// Field sorter
type FieldSorter []q.Queen

func (f FieldSorter) Len() int      { return len(f) }
func (f FieldSorter) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f FieldSorter) Less(i, j int) bool {
	return f[i].GetY() < f[j].GetY() || f[i].GetY() == f[j].GetY() && f[i].GetX() < f[j].GetX()
}
