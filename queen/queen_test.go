package queen

import (
	"testing"
	"math"
)

func TestRandomInit(t *testing.T) {
	for i := 0; i < 5; i++ {

		testSize := int(math.Exp2(float64(i)))

		testQueen := GenerateRandom(testSize)
		if (testQueen.x > testSize || testQueen.x < 0) {
			t.Errorf("Queen %d failed, mismatching Position x", testSize)
			t.Fail()
		}
	}

}

func TestEquals(t *testing.T){
	q1 := Queen{x: 4, y: 1}
	q3 := Queen{x: 4, y: 1}
	q2 := Queen{x: 4, y: 5}

	exp := false
	res := q1.Equals(q2)

	if(exp != res){
		t.Fail()
	}
	exp = true
	res = q1.Equals(q3)

	if(exp != res){
		t.Fail()
	}

}

func TestCaptures(t *testing.T) {

	q1 := Queen{x: 4, y: 1}
	q2 := Queen{x: 4, y: 5}
	q3 := Queen{x: 2, y: 2}
	q4 := Queen{x: 9, y: 2}
	q5 := Queen{x: 7, y: 4}
	q6 := Queen{x: 3, y: 8}

	//Not tested 2 3, 2 5, 3 5, 4 5

	//vertical
	exp := true
	res := q1.Captures(q2)
	if (exp != res) {
		t.Error()
	}
	res = q2.Captures(q1)
	if (exp != res) {
		t.Error()
	}
	exp = false
	res = q1.Captures(q3)
	if (exp != res) {
		t.Error()
	}
	res = q3.Captures(q1)
	if (exp != res) {
		t.Error()
	}
	//horizontal
	exp = true
	res = q3.Captures(q4)
	if (exp != res) {
		t.Error()
	}
	res = q4.Captures(q3)
	if (exp != res) {
		t.Error()
	}
	exp = false
	res = q4.Captures(q2)
	if (exp != res) {
		t.Error()
	}
	res = q2.Captures(q4)
	if (exp != res) {
		t.Error()
	}
	//cross right
	exp = true
	res = q1.Captures(q5)
	if (exp != res) {
		t.Error()
	}
	res = q5.Captures(q1)
	if (exp != res) {
		t.Error()
	}
	exp = false
	res = q1.Captures(q4)
	if (exp != res) {
		t.Error()
	}
	res = q4.Captures(q1)
	if (exp != res) {
		t.Error()
	}

	//cross left
	exp = true
	res = q5.Captures(q6)
	if (exp != res) {
		t.Error()
	}
	res = q6.Captures(q5)
	if (exp != res) {
		t.Error()
	}
	exp = false
	res = q2.Captures(q6)
	if (exp != res) {
		t.Error()
	}
	res = q6.Captures(q2)
	if (exp != res) {
		t.Error()
	}
}
