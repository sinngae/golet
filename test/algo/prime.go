package algo

import (
	"fmt"
	"math"
)

func PrimeNumGen() {
	primeNums := []int64{2}
	var rates []float64
	for val := int64(3); ; val++ {
		isPrime := true
		sqrtVal := math.Sqrt(float64(val))
		for _, primeItem := range primeNums {
			if val%primeItem == 0 {
				isPrime = false
				break
			}
			if float64(primeItem) > sqrtVal {
				break
			}
		}
		if isPrime {
			primeNums = append(primeNums, val)
		}
		if val%1000 == 0 {
			rate := float64(len(primeNums)) / float64(val)
			//fmt.Printf("val: %d, prime rate:%.4f\n", val, rate)
			//time.Sleep(1 * time.Second)
			rates = append(rates, rate)
			if len(rates)%1000 == 0 {
				fmt.Printf("%v\n", rates)
			}
		}
	}
}
