package VehicleDataGenerator

import (
	"math/rand"
	"time"
)

func generateData() ([]interface{}, error) {

	var data []interface{}
	rand.Seed(time.Now().UnixNano())

	// Reading time between events duration
	timeBetweenEvents, err := getDurationFromEnv("TIME_BETWEEN_EVENTS")
	if err != nil {
		return data, err
	}

	// Generating connect data
	connectJson := ConnectJson{
		Index:     0,
		Type:      "connect_vehicle",
		Timestamp: getTimestamp(timeBetweenEvents),
		Vin:       "car1",
	}
	data = append(data, connectJson)

	connectJson = ConnectJson{
		Index:     0,
		Type:      "connect_vehicle",
		Timestamp: getTimestamp(timeBetweenEvents),
		Vin:       "car2",
	}
	data = append(data, connectJson)

	// Generating update data
	var previousJsonCar1 UpdateJson
	var previousJsonCar2 UpdateJson
	var updateJson UpdateJson
	for i := 0; i < 10; i++ {
		if previousJsonCar1 == (UpdateJson{}) {

			// Car1
			locationJson := GpsLocation{
				Latitude:  48.154024,
				Longitude: 17.071208,
			}
			vehicleJson := Vehicle{
				Vin:                     "car1",
				FrontLidarDistance:      100.1,
				FrontUltrasonicDistance: 100.12,
				RearUltrasonicDistance:  10.19,
				WheelSpeed:              0.09,
				GpsLocation:             locationJson,
				GpsSpeed:                0.1,
				GpsDirection:            0,
				MagnetometerDirection:   0.01,
			}
			updateJson = UpdateJson{
				Index:     1,
				Type:      "update_vehicle",
				Timestamp: getTimestamp(timeBetweenEvents),
				Vehicle:   vehicleJson,
			}
			previousJsonCar1 = updateJson
			data = append(data, updateJson)

			// Car 2
			locationJson = GpsLocation{
				Latitude:  48.154014,
				Longitude: 17.071208,
			}
			vehicleJson = Vehicle{
				Vin:                     "car2",
				FrontLidarDistance:      10.2,
				FrontUltrasonicDistance: 10.19,
				RearUltrasonicDistance:  100.11,
				WheelSpeed:              0.09,
				GpsLocation:             locationJson,
				GpsSpeed:                0.1,
				GpsDirection:            0,
				MagnetometerDirection:   0.01,
			}
			updateJson = UpdateJson{
				Index:     1,
				Type:      "update_vehicle",
				Timestamp: getTimestamp(timeBetweenEvents),
				Vehicle:   vehicleJson,
			}
			previousJsonCar2 = updateJson
			data = append(data, updateJson)
		} else {

			speedIncrement := 0.1 + rand.Float64()*(1-0.1)
			distanceIncrement := -0.1 + rand.Float64()*(0.1+0.1)

			// Car1
			locationJson := GpsLocation{
				Latitude:  previousJsonCar1.Vehicle.GpsLocation.Latitude + speedIncrement*0.00001,
				Longitude: 17.071208,
			}
			vehicleJson := Vehicle{
				Vin:                     "car1",
				FrontLidarDistance:      100.1,
				FrontUltrasonicDistance: 100.12,
				RearUltrasonicDistance:  previousJsonCar1.Vehicle.RearUltrasonicDistance + distanceIncrement,
				WheelSpeed:              previousJsonCar1.Vehicle.WheelSpeed + speedIncrement,
				GpsLocation:             locationJson,
				GpsSpeed:                previousJsonCar1.Vehicle.GpsSpeed + speedIncrement,
				GpsDirection:            0,
				MagnetometerDirection:   0.01,
			}
			updateJson = UpdateJson{
				Index:     1 + i,
				Type:      "update_vehicle",
				Timestamp: getTimestamp(timeBetweenEvents),
				Vehicle:   vehicleJson,
			}
			previousJsonCar1 = updateJson
			data = append(data, updateJson)

			// Car 2
			locationJson = GpsLocation{
				Latitude:  previousJsonCar2.Vehicle.GpsLocation.Latitude + speedIncrement*0.00001,
				Longitude: 17.071208,
			}
			vehicleJson = Vehicle{
				Vin:                     "car2",
				FrontLidarDistance:      previousJsonCar2.Vehicle.FrontLidarDistance + distanceIncrement,
				FrontUltrasonicDistance: previousJsonCar2.Vehicle.FrontUltrasonicDistance + distanceIncrement,
				RearUltrasonicDistance:  100.11,
				WheelSpeed:              previousJsonCar2.Vehicle.WheelSpeed + speedIncrement,
				GpsLocation:             locationJson,
				GpsSpeed:                previousJsonCar2.Vehicle.GpsSpeed + speedIncrement,
				GpsDirection:            0,
				MagnetometerDirection:   0.01,
			}
			updateJson = UpdateJson{
				Index:     1 + i,
				Type:      "update_vehicle",
				Timestamp: getTimestamp(timeBetweenEvents),
				Vehicle:   vehicleJson,
			}
			previousJsonCar2 = updateJson
			data = append(data, updateJson)
		}
	}

	return data, nil
}
