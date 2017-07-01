package field

import (
	"testing"
)

func TestRandomInit(t *testing.T)  {
	testField := GenerateRandom(10)
	testField.Print()
	testField.GetFitness()
	testField2 := GenerateRandom(10)
	recombined := testField2.Recombine(testField)
	recombined.Print()
}