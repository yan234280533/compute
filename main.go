package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func main() {

	var input []float64
	/*for i := 0; i < 1000; i++ {
		input = append(input, float64(i))
	}*/
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		value := rand.Intn(3000)
		input = append(input, float64(value))
	}

	sort.Float64s(input)

	s1 := sum1(input, 1.0)
	fmt.Printf("s1: %.4f]\n", s1)

	var minValue = math.MaxFloat64
	var minIndex = -1

	for i := 0; i < len(input); i++ {
		s2 := sum2(input, 1.0, 0.75, input[i])
		fmt.Printf("s2: [%d,v:%.4f,%.4f]\n", i, input[i], s2)

		if s2 <= minValue {
			minIndex = i
			minValue = s2
		}
	}

	fmt.Printf("min: [%d,%.4f]\n", minIndex, minValue)

	return
}

func sum1(input []float64, uintPrice float64) float64 {
	var sum float64

	for _, v := range input {
		sum += uintPrice * v
	}

	return sum
}

func sum2(input []float64, uintPrice float64, uintPriceMonthly float64, monthlyScale float64) float64 {

	var sum = float64(len(input)) * uintPriceMonthly * monthlyScale

	for _, v := range input {
		if v <= monthlyScale {
			continue
		}

		sum += uintPrice * (v - monthlyScale)
	}

	return sum
}
