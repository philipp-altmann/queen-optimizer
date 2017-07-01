package field

import (
	q "github.com/philipp-altmann/QueenOptimizer/queen"
	"fmt"
	"strings"
	"sort"
	"time"
	"math/rand"
	"math"
)

type Field struct {
	queens  [] q.Queen
	size    int
	fitness int
}

func Generate(queens []q.Queen, size int)(field Field)  {
	return Field{queens: queens, size: size, fitness: -1}
}

func GenerateRandom(fieldSize int)(field Field)  {
	var queens [] q.Queen
	for i := 0; i < fieldSize; i++ {
		queens = append(queens, q.GenerateRandom(fieldSize))
	}
	return Field{queens: queens, size: fieldSize, fitness: -1}
}

func (field *Field)Print() {

	var fieldArray = make([][]bool, field.size)//[field.size][field.size] bool
	for i := range fieldArray {
		fieldArray[i] = make([]bool, field.size)
	}
	for i := 0; i < field.size; i++ {
		fieldArray[field.queens[i].GetX()][field.queens[i].GetY()] = true
	}
	fmt.Print("\n+" + strings.Repeat("–––+", len(fieldArray)) + "\n")

	for i := 0; i < len(fieldArray); i++ {
		fmt.Print("|")
		for j := 0; j < len(fieldArray[i]); j++ {
			if (fieldArray[i][j]) {
				fmt.Print(" X |")
			} else {
				fmt.Print("   |")
			}
		}
		fmt.Print("\n+" + strings.Repeat("–––+", len(fieldArray[i])) + "\n")
	}
}

func (field *Field)Evaluate() () {
	fitness := 0
	equal := 0
	for i := 0; i < field.size; i++ {
		for j := i + 1; j < field.size; j++ {
			/*fmt.Printf("Comparing Queen at (%d/%d) with queen at (%d/%d)... \n",
				field.queens[i].positionX, field.queens[i].positionY,
				field.queens[j].positionX, field.queens[j].positionY )*/
			if field.queens[i].Captures(field.queens[j]) {
				fitness++
				if (field.queens[i].Equals(field.queens[j])) {
					equal ++;
				}
				//fmt.Printf("| Captured\n")
				/*fmt.Printf("Queen at (%d/%d) captures queen at (%d/%d)\n",
					field.queens[i].positionX, field.queens[i].positionY,
					field.queens[j].positionX, field.queens[j].positionY )*/
			}
		}
	}
	if(equal != 0){
		//const worstFitness = 120 //todo calc

		//fitness = worstFitness
		fitness = math.MaxInt64

	}
	field.fitness = fitness
	//fmt.Printf("Fittness: %d, %d Equal\n", fitness, equal)
}

func (field *Field)Approximate(approximationPool []Field)()  {
	var similarity []float64
	for i:=0; i<len(approximationPool); i++ {
		similarity = append(similarity, 1 - field.Compare(approximationPool[i]))
	}
	var similaritySum float64 = 0
	for i := range similarity {
		similaritySum += similarity[i]
	}

	var approximatefitness float64 = 0
	for i:=0; i<len(approximationPool); i++ {
		approximatefitness += (similarity[i]/similaritySum)*float64(approximationPool[i].fitness)
	}

	field.fitness = int(math.Floor(approximatefitness))

}

func (field *Field)Compare(to Field)(difference float64)  {
	sort.Sort(FieldSorter(field.queens))
	sort.Sort(FieldSorter(to.queens))
	//fmt.Printf("Queens: %v\n",field.queens)
	//fmt.Printf("Queens: %v\n",to.queens)
	for i:=0; i<field.size; i++ {
		dx:= math.Abs(float64(field.queens[i].GetX()-to.queens[i].GetX()))/(float64(field.size-1))
		dy:= math.Abs(float64(field.queens[i].GetY()-to.queens[i].GetY()))/(float64(field.size-1))
		//fmt.Printf("Queens at %d, Deltas: X: %f, Y: %f\n",i, dx, dy )
		difference += dx
		difference += dy
	}
	difference /= float64(field.size*2)
	if(difference >= 1.0){
		fmt.Print("Found same field")
	}
	return
}



func (field *Field)GetFitness()(fitness int)  {
	return field.fitness
}

func (field1 *Field)Recombine(field2 Field) (field Field) {
	sort.Sort(FieldSorter(field1.queens))
	sort.Sort(FieldSorter(field2.queens))

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	cut := r.Intn(field1.size)
	var queens [] q.Queen
	for i := 0; i < cut; i++ {
		queens = append(queens, q.Generate(field1.queens[i].GetX(), field1.queens[i].GetY()))
	}
	for j := cut; j < field1.size; j++ {
		queens = append(queens, q.Generate(field2.queens[j].GetX(), field2.queens[j].GetY()))
	}
	return Field{queens: queens, size: field1.size}
}

func (field *Field)Mutate()()  {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var queen = r.Intn(field.size)

	field.queens[queen].Mutate(field.size)
}

// Field sorter
type FieldSorter []q.Queen

func (f FieldSorter) Len() int           { return len(f) }
func (f FieldSorter) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f FieldSorter) Less(i, j int) bool { return f[i].GetY() < f[j].GetY() || f[i].GetY() == f[j].GetY() && f[i].GetX() < f[j].GetX() }
