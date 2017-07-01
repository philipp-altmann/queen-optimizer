package queen

import (
	"math"
	"time"
	"math/rand"
)

type Queen struct {
	x int	//Position X
	y int	//Position Y
}

func Generate(x int, y int)(queen Queen)  {
	return Queen{x: x, y: y}
}

func GenerateRandom(fieldSize int)(randomQueen Queen)  {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomQueen = Queen{
		x: r.Intn(fieldSize),
		y: r.Intn(fieldSize),
	}
	return
}

func (q* Queen)GetX()(x int)  {
	return q.x
}

func (q* Queen)GetY()(y int)  {
	return q.y
}

func (q* Queen)Equals(queen Queen) (c bool) {
	return (q.x==queen.x && q.y ==queen.y)
}
func (q* Queen)Captures(queen Queen) (c bool) {
	if (q.x == queen.x) {
		return true
	} else if (q.y == queen.y) {
		return true
	} else if (math.Abs(float64(q.x-queen.x)) == math.Abs(float64(q.y-queen.y)) ) {
		return true
	} else {
		return false
	}
}

func (queen *Queen)Mutate(fieldSize int)()  {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var posType = r.Intn(2)
	var mutation = r.Intn(5) - 2

	if (posType == 1) {
		//fmt.Printf("Mutation: %d",mutation)
		var newPos = queen.x + mutation
		if (newPos >= fieldSize || newPos < 0) {
			//DO NOTHING
			queen.x = r.Intn(fieldSize)
		} else {
			queen.x += mutation
		}
	} else {
		var newPos = queen.y + mutation
		if (newPos >= fieldSize || newPos < 0) {
			//DO NOTHING
			queen.y = r.Intn(fieldSize)

		} else {
			queen.y += mutation
		}
	}

}