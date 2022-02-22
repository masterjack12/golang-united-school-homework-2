package square

import (
	"fmt"
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type sum int

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum sum) float64 {

	var square float64

	if sidesNum != SidesTriangle && sidesNum != SidesSquare && sidesNum != SidesCircle {
		square = 0
	} else if sidesNum == SidesTriangle {
		square = math.Pow(sideLen, float64(2)) * math.Sqrt(3) / float64(4)
	} else if sidesNum == SidesSquare {
		square = math.Pow(sideLen, sideLen)
	} else if sidesNum == SidesCircle {
		square = math.Pi * math.Pow(sideLen, sideLen)
	}

	fmt.Println(square)
	return square

}
