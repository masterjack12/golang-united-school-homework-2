package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type sNumber int

const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum sNumber) float64 {

	var square float64

	if sidesNum != SidesTriangle && sidesNum != SidesSquare && sidesNum != SidesCircle {
		square = 0
	} else if sidesNum == SidesTriangle {
		square = (sideLen * sideLen) * math.Sqrt(3) / float64(4)
	} else if sidesNum == SidesSquare {
		square = sideLen * sideLen
	} else if sidesNum == SidesCircle {
		square = math.Pi * (sideLen * sideLen)
	}

	return square

}
