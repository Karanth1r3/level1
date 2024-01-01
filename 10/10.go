package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	arr := []float64{-25.4, -15.5, 5., -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 158.5}
	groupTemps(arr)
}

// assigns temperatures from input data to map keys according to rules from the example
func groupTemps(tmps []float64) map[float64][]float64 {
	m := make(map[float64][]float64)

	slices.Sort(tmps) // for getting min & max at polars of the slice

	min := roundTemp(tmps[0], false)          // getting minimum group key
	max := roundTemp(tmps[len(tmps)-1], true) // max group key
	//fmt.Println(min, max)

	mn := (min)
	mx := (max)

	c := getRange(mn, mx)
	//fmt.Println(c)
	for _, elem := range tmps {
		//	slc := make([]float64, 0)
		for _, scale := range c {
			//fmt.Println(elem, scale)
			if elem < 0 && elem+10 > scale && elem < scale {
				m[scale] = append(m[scale], elem)
				break
			} else if elem > 0 && elem-10 < scale && elem > scale {
				m[scale] = append(m[scale], elem)
				break
			}
		}
	}
	fmt.Println(m)
	return m
}

// generate range for given data (elements in range will be different by 10)
func getRange(min, max float64) []float64 {
	i := 0
	result := make([]float64, 1)

	for min < max {
		result = append(result, min)
		min += 10
		i++
	}
	result = append(result, max)
	return result
}

// get min or max value rounded to 10, ceil or floor (for min or max of the range) depends on bool argument
func roundTemp(val float64, isMax bool) float64 {
	dgts := 0 // number of digits in number to recreate it later
	if val > 0 {
		for val > 10 {
			val /= 10
			dgts++
		}
	} else {
		for val < -10 {
			val /= 10
			dgts++
		}
	}
	if isMax { // if need to find max range value - floor it, otherwise - ceil
		val = math.Floor(val)
	} else {
		val = math.Ceil(val)
	}
	mul := math.Pow(10, float64(dgts)) //
	return val * mul
}
