package field

import (
	"fmt"
	"strings"
	"testing"

	q "github.com/philipp-altmann/QueenOptimizer/queen"
)

const testFieldSize = 8

func TestInit(t *testing.T) {
	fmt.Print("Testing Initialization ...\n")

	testField := GenerateRandom(testFieldSize)
	if testField.size != testFieldSize {
		t.Fail()
	}
	if len(testField.queens) != testFieldSize {
		t.Fail()
	}

	fmt.Print("Random Tests succeeded.\n\n")
}

func TestPrint(t *testing.T) {
	fmt.Print("Testing Printing ...\n")

	testField := GenerateRandom(testFieldSize)
	printable := testField.ToPrintable()

	//Remove styling
	printable = strings.Replace(printable, "–", "", -1) //strings.Trim(printable, " ")
	printable = strings.Replace(printable, "+", "", -1)
	printable = strings.Replace(printable, "|", "", -1)
	printable = strings.Replace(printable, " X ", "X", -1)
	printable = strings.Replace(printable, "   ", " ", -1)
	printable = strings.Replace(printable, "\n", "", -1)

	if len(printable) != testFieldSize*testFieldSize {
		t.Fail()
	}

	for i := 0; i < testFieldSize; i++ {
		x := testField.queens[i].GetX()
		y := testField.queens[i].GetY()
		pos := x + y*testFieldSize
		if printable[pos] != 'X' {
			t.Fail()
		}
	}

	fmt.Print("Print Tests succeeded.\n\n")
}

func TestEvaluation(t *testing.T) {
	fmt.Print("Testing Evaluation ...\n")

	testField := GenerateRandom(testFieldSize)
	testField.Evaluate()
	const maxFitness = testFieldSize * testFieldSize
	if testField.GetFitness() < 0 {
		t.Fail()
	}
	if testField.GetFitness() > maxFitness {
		t.Log("Fitness too large")
		t.Fail()
	}

	var queens []q.Queen
	for i := 0; i < testFieldSize; i++ {
		queens = append(queens, q.Generate(0, 0))
	}
	worstField := Generate(queens, testFieldSize)
	worstField.Evaluate()
	if worstField.GetFitness() != maxFitness {
		t.Fail()
	}

	fmt.Print("Evaluation Tests succeeded.\n\n")
}

func TestDistance(t *testing.T) {
	fmt.Print("Testing Distance ...\n")

	testField := GenerateRandom(testFieldSize)
	compareField := GenerateRandom(testFieldSize)
	if testField.Distance(compareField) < 0 {
		t.Fail()
	}
	if testField.Distance(testField) != 0 {
		t.Fail()
	}

	fmt.Print("Distance Tests succeeded.\n\n")
}

func TestCompare(t *testing.T) {
	fmt.Print("Testing Compare ...\n")

	testField := GenerateRandom(testFieldSize)
	compareField := GenerateRandom(testFieldSize)

	testField.Compare(compareField)
	if testField.Compare(compareField) < 0 {
		t.Log("Too Small")
		t.Fail()
	}
	if testField.Compare(compareField) > 1 {
		t.Log("Too Large")
		t.Fail()
	}
	if testField.Compare(testField) != 0 {
		t.Log("Not Equal")
		t.Fail()
	}

	fmt.Print("Compare Tests succeeded.\n\n")

}

func TestMutation(t *testing.T) {
	fmt.Print("Testing Mutation ...\n")
	testField := GenerateRandom(testFieldSize)
	//Deep Copy Field
	var queens []q.Queen
	for i := 0; i < testFieldSize; i++ {
		queens = append(queens, q.Generate(testField.queens[i].GetX(), testField.queens[i].GetY()))
	}
	mutated := Generate(queens, testFieldSize)

	mutated.Mutate()

	if testField.Distance(mutated) == 0 {
		t.Fail()
	}
	fmt.Print("Mutation Tests succeeded.\n\n")

}

func TestRecombination(t *testing.T) {
	fmt.Print("Testing Recombination ...\n")

	testField := GenerateRandom(testFieldSize)
	testField2 := GenerateRandom(testFieldSize)
	for testField.Distance(testField2) == 0 {
		testField2 = GenerateRandom(testFieldSize)
	}
	recombined := testField2.Recombine(testField)
	if recombined.size != testFieldSize {
		t.Fail()
	}
	if len(recombined.queens) != testFieldSize {
		t.Fail()
	}
	if recombined.Distance(testField) == 0 {
		t.Fail()
		t.Log("Resulting field equals input")
	}
	if recombined.Distance(testField2) == 0 {
		t.Fail()
		t.Log("Resulting field equals input")
	}
	fmt.Print("Recombination Tests succeeded.\n\n")
}
