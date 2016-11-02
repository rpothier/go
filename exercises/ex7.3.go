package main

import (
	"fmt"
	"math"
)

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegativeSqrt(f)
	}
	return math.Sqrt(f), nil
}

type ErrNegativeSqrt float64

//and make it an error by giving it a

func (e ErrNegativeSqrt) Error() string{
    	return "Error - Cannot get sqare rt of neg num"
}



func main() {
	fmt.Println(sqrt(2.0))
	fmt.Println(sqrt(-2.0))
}
