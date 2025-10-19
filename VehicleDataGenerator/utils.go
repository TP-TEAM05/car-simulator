package VehicleDataGenerator

import (
	"os"
	"strconv"
	"strings"
	"time"
)

var TIMESTAMP time.Time
var TimestampInitialized = false

func getTimestamp(timeBetweenEvents time.Duration) string {
	if TimestampInitialized {
		TIMESTAMP = TIMESTAMP.Add(timeBetweenEvents)
	} else {
		TIMESTAMP = time.Now().UTC()
		TimestampInitialized = true
	}
	tStr := TIMESTAMP.String()
	tStr = strings.Replace(tStr, " ", "T", 1)
	tStr = tStr[:23]
	tStr = tStr + "Z"
	return tStr
}

func GetDurationFromEnv(key string) (time.Duration, error) {
	var result time.Duration
	result, err := time.ParseDuration(os.Getenv(key))
	if err != nil {
		return result, err
	}
	return result, nil
}

func GetFloatFromEnv(key string) (float64, error) {
	result, err := strconv.ParseFloat(os.Getenv(key), 64)
	if err != nil {
		return result, err
	}
	return result, nil
}
