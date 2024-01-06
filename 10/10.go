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

	min := roundTemp(tmps[0], 10)           // getting minimum group key
	max := roundTemp(tmps[len(tmps)-1], 10) // max group key
	fmt.Println(min, max)

	mn := (min)
	mx := (max)

	c := getRange(mn, mx) // can be changed to allow any step in argument actually (but not today) currently works for 10 only
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

// get min or max value rounded to step, ceil or floor (for min or max of the range) depends on bool argument
func roundTemp(val float64, step int) float64 {
	var temp int = int(math.Abs(val / float64(step)))
	fmt.Println(temp)
	if val > 0 {
		val = float64(step*temp + step)
	} else {
		val = float64(-step*temp - step)
	}

	return val
}
