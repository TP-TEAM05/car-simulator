package VehicleDataGenerator

import (
	"math/rand"
	"time"
)

func generateData() ([]interface{}, error) {

	var data []interface{}
	rand.Seed(time.Now().UnixNano())

	// Reading time between events duration
	timeBetweenEvents, err := GetDurationFromEnv("TIME_BETWEEN_EVENTS")
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

	// Going in a straight line to North
	for i := 0; i < 40; i++ {
		if previousJsonCar1 == (UpdateJson{}) {

			// Car1
			locationJson := GpsLocation{
				Latitude:  48.154054,
				Longitude: 17.071268,
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
				Timestamp: "TimestampToReplace",
				Vehicle:   vehicleJson,
			}
			previousJsonCar1 = updateJson

			updateJson.Index = 0

			data = append(data, updateJson)

			// Car 2
			locationJson = GpsLocation{
				Latitude:  48.154014,
				Longitude: 17.071268,
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
				Timestamp: "TimestampToReplace",
				Vehicle:   vehicleJson,
			}
			previousJsonCar2 = updateJson

			updateJson.Index = 0

			data = append(data, updateJson)
		} else {

			speedIncrement := 0.1 + rand.Float64()*(1-0.1)
			distanceIncrement := -0.1 + rand.Float64()*(0.1+0.1)

			// Car1
			locationJson := GpsLocation{
				Latitude:  previousJsonCar1.Vehicle.GpsLocation.Latitude + previousJsonCar1.Vehicle.GpsSpeed*0.0000005,
				Longitude: previousJsonCar1.Vehicle.GpsLocation.Longitude,
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
				Index:     previousJsonCar1.Index + 1,
				Type:      "update_vehicle",
				Timestamp: "TimestampToReplace",
				Vehicle:   vehicleJson,
			}
			previousJsonCar1 = updateJson

			updateJson.Index = 0

			data = append(data, updateJson)

			// Car 2
			locationJson = GpsLocation{
				Latitude:  previousJsonCar2.Vehicle.GpsLocation.Latitude + previousJsonCar2.Vehicle.GpsSpeed*0.0000005,
				Longitude: previousJsonCar2.Vehicle.GpsLocation.Longitude,
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
				Index:     previousJsonCar2.Index + 1,
				Type:      "update_vehicle",
				Timestamp: "TimestampToReplace",
				Vehicle:   vehicleJson,
			}
			previousJsonCar2 = updateJson

			updateJson.Index = 0

			data = append(data, updateJson)
		}
	}

	previousJsonCar1.Vehicle.MagnetometerDirection = 360
	previousJsonCar1.Vehicle.GpsDirection = 360
	previousJsonCar2.Vehicle.MagnetometerDirection = 360
	previousJsonCar2.Vehicle.GpsDirection = 360

	// Turning 90 degrees to the West
	for i := 0; i < 9; i++ {
		distanceIncrement := -0.1 + rand.Float64()*(0.1+0.1)

		// Car1
		locationJson := GpsLocation{
			Latitude:  previousJsonCar1.Vehicle.GpsLocation.Latitude + 0.000001,
			Longitude: previousJsonCar1.Vehicle.GpsLocation.Longitude - 0.000001,
		}
		vehicleJson := Vehicle{
			Vin:                     "car1",
			FrontLidarDistance:      100.1,
			FrontUltrasonicDistance: 100.12,
			RearUltrasonicDistance:  previousJsonCar1.Vehicle.RearUltrasonicDistance + distanceIncrement,
			WheelSpeed:              previousJsonCar1.Vehicle.WheelSpeed,
			GpsLocation:             locationJson,
			GpsSpeed:                previousJsonCar1.Vehicle.GpsSpeed,
			GpsDirection:            previousJsonCar1.Vehicle.GpsDirection - 2.5,
			MagnetometerDirection:   previousJsonCar1.Vehicle.MagnetometerDirection - 2.5,
		}
		updateJson = UpdateJson{
			Index:     previousJsonCar1.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar1 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)

		// Car 2
		locationJson = GpsLocation{
			Latitude:  previousJsonCar2.Vehicle.GpsLocation.Latitude + 0.000001,
			Longitude: previousJsonCar2.Vehicle.GpsLocation.Longitude - 0.000001,
		}
		vehicleJson = Vehicle{
			Vin:                     "car2",
			FrontLidarDistance:      previousJsonCar2.Vehicle.FrontLidarDistance + distanceIncrement,
			FrontUltrasonicDistance: previousJsonCar2.Vehicle.FrontUltrasonicDistance + distanceIncrement,
			RearUltrasonicDistance:  100.11,
			WheelSpeed:              previousJsonCar2.Vehicle.WheelSpeed,
			GpsLocation:             locationJson,
			GpsSpeed:                previousJsonCar2.Vehicle.GpsSpeed,
			GpsDirection:            previousJsonCar2.Vehicle.GpsDirection - 10,
			MagnetometerDirection:   previousJsonCar2.Vehicle.MagnetometerDirection - 10,
		}
		updateJson = UpdateJson{
			Index:     previousJsonCar2.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar2 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)
	}

	// Turning 90 degrees to the South
	for i := 0; i < 9; i++ {
		distanceIncrement := -0.1 + rand.Float64()*(0.1+0.1)

		// Car1
		locationJson := GpsLocation{
			Latitude:  previousJsonCar1.Vehicle.GpsLocation.Latitude - 0.000001,
			Longitude: previousJsonCar1.Vehicle.GpsLocation.Longitude - 0.000001,
		}
		vehicleJson := Vehicle{
			Vin:                     "car1",
			FrontLidarDistance:      100.1,
			FrontUltrasonicDistance: 100.12,
			RearUltrasonicDistance:  previousJsonCar1.Vehicle.RearUltrasonicDistance + distanceIncrement,
			WheelSpeed:              previousJsonCar1.Vehicle.WheelSpeed,
			GpsLocation:             locationJson,
			GpsSpeed:                previousJsonCar1.Vehicle.GpsSpeed,
			GpsDirection:            previousJsonCar1.Vehicle.GpsDirection - 10,
			MagnetometerDirection:   previousJsonCar1.Vehicle.MagnetometerDirection - 10,
		}
		updateJson = UpdateJson{
			Index:     previousJsonCar1.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar1 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)

		// Car 2
		locationJson = GpsLocation{
			Latitude:  previousJsonCar2.Vehicle.GpsLocation.Latitude - 0.000001,
			Longitude: previousJsonCar2.Vehicle.GpsLocation.Longitude - 0.000001,
		}
		vehicleJson = Vehicle{
			Vin:                     "car2",
			FrontLidarDistance:      previousJsonCar2.Vehicle.FrontLidarDistance + distanceIncrement,
			FrontUltrasonicDistance: previousJsonCar2.Vehicle.FrontUltrasonicDistance + distanceIncrement,
			RearUltrasonicDistance:  100.11,
			WheelSpeed:              previousJsonCar2.Vehicle.WheelSpeed,
			GpsLocation:             locationJson,
			GpsSpeed:                previousJsonCar2.Vehicle.GpsSpeed,
			GpsDirection:            previousJsonCar2.Vehicle.GpsDirection - 2.5,
			MagnetometerDirection:   previousJsonCar2.Vehicle.MagnetometerDirection - 2.5,
		}
		updateJson = UpdateJson{
			Index:     previousJsonCar2.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar2 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)
	}

	// Going in a straight line to South
	for i := 0; i < 40; i++ {

		speedIncrement := -0.1 + rand.Float64()*(0.1+0.1)
		distanceIncrement := -0.1 + rand.Float64()*(0.1+0.1)

		// Car1
		locationJson := GpsLocation{
			Latitude:  previousJsonCar1.Vehicle.GpsLocation.Latitude - previousJsonCar1.Vehicle.GpsSpeed*0.00000025,
			Longitude: previousJsonCar1.Vehicle.GpsLocation.Longitude,
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
			Index:     previousJsonCar1.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar1 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)

		// Car 2
		locationJson = GpsLocation{
			Latitude:  previousJsonCar2.Vehicle.GpsLocation.Latitude - previousJsonCar2.Vehicle.GpsSpeed*0.00000025,
			Longitude: previousJsonCar2.Vehicle.GpsLocation.Longitude,
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
			Index:     previousJsonCar2.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar2 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)
	}

	// Turning 90 degrees to the East
	for i := 0; i < 9; i++ {
		distanceIncrement := -0.1 + rand.Float64()*(0.1+0.1)

		// Car1
		locationJson := GpsLocation{
			Latitude:  previousJsonCar1.Vehicle.GpsLocation.Latitude - 0.000001,
			Longitude: previousJsonCar1.Vehicle.GpsLocation.Longitude + 0.000001,
		}
		vehicleJson := Vehicle{
			Vin:                     "car1",
			FrontLidarDistance:      100.1,
			FrontUltrasonicDistance: 100.12,
			RearUltrasonicDistance:  previousJsonCar1.Vehicle.RearUltrasonicDistance + distanceIncrement,
			WheelSpeed:              previousJsonCar1.Vehicle.WheelSpeed,
			GpsLocation:             locationJson,
			GpsSpeed:                previousJsonCar1.Vehicle.GpsSpeed,
			GpsDirection:            previousJsonCar1.Vehicle.GpsDirection - 10,
			MagnetometerDirection:   previousJsonCar1.Vehicle.MagnetometerDirection - 10,
		}
		updateJson = UpdateJson{
			Index:     previousJsonCar1.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar1 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)

		// Car 2
		locationJson = GpsLocation{
			Latitude:  previousJsonCar2.Vehicle.GpsLocation.Latitude - 0.000001,
			Longitude: previousJsonCar2.Vehicle.GpsLocation.Longitude + 0.000001,
		}
		vehicleJson = Vehicle{
			Vin:                     "car2",
			FrontLidarDistance:      previousJsonCar2.Vehicle.FrontLidarDistance + distanceIncrement,
			FrontUltrasonicDistance: previousJsonCar2.Vehicle.FrontUltrasonicDistance + distanceIncrement,
			RearUltrasonicDistance:  100.11,
			WheelSpeed:              previousJsonCar2.Vehicle.WheelSpeed,
			GpsLocation:             locationJson,
			GpsSpeed:                previousJsonCar2.Vehicle.GpsSpeed,
			GpsDirection:            previousJsonCar2.Vehicle.GpsDirection - 2.5,
			MagnetometerDirection:   previousJsonCar2.Vehicle.MagnetometerDirection - 2.5,
		}
		updateJson = UpdateJson{
			Index:     previousJsonCar2.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar2 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)
	}

	// Turning 90 degrees to the North
	for i := 0; i < 9; i++ {
		distanceIncrement := -0.1 + rand.Float64()*(0.1+0.1)

		// Car1
		locationJson := GpsLocation{
			Latitude:  previousJsonCar1.Vehicle.GpsLocation.Latitude + 0.000001,
			Longitude: previousJsonCar1.Vehicle.GpsLocation.Longitude + 0.000001,
		}
		vehicleJson := Vehicle{
			Vin:                     "car1",
			FrontLidarDistance:      100.1,
			FrontUltrasonicDistance: 100.12,
			RearUltrasonicDistance:  previousJsonCar1.Vehicle.RearUltrasonicDistance + distanceIncrement,
			WheelSpeed:              previousJsonCar1.Vehicle.WheelSpeed,
			GpsLocation:             locationJson,
			GpsSpeed:                previousJsonCar1.Vehicle.GpsSpeed,
			GpsDirection:            previousJsonCar1.Vehicle.GpsDirection - 10,
			MagnetometerDirection:   previousJsonCar1.Vehicle.MagnetometerDirection - 10,
		}
		updateJson = UpdateJson{
			Index:     previousJsonCar1.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar1 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)

		// Car 2
		locationJson = GpsLocation{
			Latitude:  previousJsonCar2.Vehicle.GpsLocation.Latitude + 0.000001,
			Longitude: previousJsonCar2.Vehicle.GpsLocation.Longitude + 0.000001,
		}
		vehicleJson = Vehicle{
			Vin:                     "car2",
			FrontLidarDistance:      previousJsonCar2.Vehicle.FrontLidarDistance + distanceIncrement,
			FrontUltrasonicDistance: previousJsonCar2.Vehicle.FrontUltrasonicDistance + distanceIncrement,
			RearUltrasonicDistance:  100.11,
			WheelSpeed:              previousJsonCar2.Vehicle.WheelSpeed,
			GpsLocation:             locationJson,
			GpsSpeed:                previousJsonCar2.Vehicle.GpsSpeed,
			GpsDirection:            previousJsonCar2.Vehicle.GpsDirection - 10,
			MagnetometerDirection:   previousJsonCar2.Vehicle.MagnetometerDirection - 10,
		}
		updateJson = UpdateJson{
			Index:     previousJsonCar2.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar2 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)
	}

	return data, nil
}
