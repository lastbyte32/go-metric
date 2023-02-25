package utils

import (
	"fmt"
	"strconv"
)

func StringToFloat64(value string) (float64, error) {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("StringToFloat64 err parse")
		return float64(0), nil
	}
	return f, nil
}

func StringToInt64(value string) (int64, error) {
	f, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("StringToInt64 err parse")
		return int64(0), err
	}
	return f, nil
}
