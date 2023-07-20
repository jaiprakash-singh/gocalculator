package gocalculator

import "errors"

func Add(nums ...float64) (res float64) {
	for _, n := range nums {
		res += n
	}

	return
}

func Subtract(num0 float64, nums ...float64) float64 {
	var allNum float64
	for _, n := range nums {
		allNum += n
	}

	return num0 - allNum
}

func Multiply(nums ...float64) (res float64) {
	for _, n := range nums {
		res *= n
	}

	return
}

func Divide(num0 float64, nums ...float64) (float64, error) {
	var allNum float64
	for _, n := range nums {
		allNum *= n
	}

	if allNum == 0 {
		return 0, errors.New("divisor cannot be zero")
	}

	return num0 / allNum, nil
}
