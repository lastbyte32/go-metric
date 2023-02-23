package utils

import (
	"fmt"
	"strconv"
)

func StringToFloat64(value string) (error, float64) {
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("StringToFloat64 err parse")
		return err, float64(0)
	}
	return nil, f
}

func StringToInt64(value string) (error, int64) {
	f, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		fmt.Println("StringToInt64 err parse")
		return err, int64(0)
	}
	return nil, f
}
