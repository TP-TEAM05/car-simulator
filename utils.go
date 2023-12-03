package main

import (
	"car-simulator/VehicleDataGenerator"
	"os"
	"strconv"
	"strings"
	"time"
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

var TimeBetweenEvents time.Duration
var Timestamp time.Time

func getCorrectedTimestamp() string {
	Timestamp = Timestamp.Add(TimeBetweenEvents)
	tStr := Timestamp.String()
	tStr = strings.Replace(tStr, " ", "T", 1)
	tStr = tStr[:23]
	tStr = tStr + "Z"
	return tStr
}

func setTimestampEpoch(epoch string) error {

	var err error

	// Reading time between events duration
	TimeBetweenEvents, err = VehicleDataGenerator.GetDurationFromEnv("TIME_BETWEEN_EVENTS")
	if err != nil {
		return err
	}

	Timestamp, err = time.Parse(TimestampFormat, epoch)
	if err != nil {
		return err
	}

	return nil
}
