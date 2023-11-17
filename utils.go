package main

import (
	"os"
	"strconv"
)

func contains(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

func GetFloatFromEnv(key string, defaultValue float32) float32 {
	result, err := strconv.ParseFloat(os.Getenv(key), 32)
	if err != nil {
		return defaultValue
	}
	return float32(result)
}
