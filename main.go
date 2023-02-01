package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func area(radius float64) (float64, error) {
	if radius < 0 {
		err := errors.New("radius must be > 0")
		return -1, err
	}
	result := math.Pi * radius * radius
	return result, nil
}

func main() {
	radius := 25.0

	result, err := area(radius)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
