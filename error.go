package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	result, err := sqrt(25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("negatif sayı bulunamadı")
	}
	return math.Sqrt(x), nil
}
