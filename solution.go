package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type intSidesNum int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

const (
	SidesTriangle intSidesNum = 3
	SidesSquare intSidesNum = 4
	SidesCircle intSidesNum = 0
)

func CalcSquare(sideLen float64, sidesNum intSidesNum) float64 {
	switch sidesNum {
	case SidesTriangle:
		return (sideLen * sideLen) / 2
	case SidesSquare:
		return sideLen * sideLen
	case SidesCircle:
		return math.Pi * (sideLen * sideLen)
	default:
		return 0
	}
}

