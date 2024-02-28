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

func generateNewData() ([]interface{}, error) {
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
	var previousJsonCar1 NewUpdateJson
	var previousJsonCar2 NewUpdateJson
	var updateJson NewUpdateJson

	// Going in a straight line to North
	for i := 0; i < 40; i++ {
		if previousJsonCar1 == (NewUpdateJson{}) {

			// Car1
			vehicleJson := NewVehicle{
				Vin:                "car1",
				Longitude:          17.071268,
				Latitude:           48.154054,
				DistanceUltrasonic: 100.12,
				DistanceLidar:      100.1,
				SpeedFrontLeft:     0.1,
				SpeedFrontRight:    0.1,
				SpeedRearLeft:      0.1,
				SpeedRearRight:     0.1,
			}
			updateJson = NewUpdateJson{
				Index:     1,
				Type:      "update_vehicle",
				Timestamp: "TimestampToReplace",
				Vehicle:   vehicleJson,
			}
			previousJsonCar1 = updateJson

			updateJson.Index = 0

			data = append(data, updateJson)

			// Car 2
			vehicleJson = NewVehicle{
				Vin:                "car2",
				Longitude:          17.071268,
				Latitude:           48.154014,
				DistanceUltrasonic: 100.19,
				DistanceLidar:      100.2,
				SpeedFrontLeft:     0.1,
				SpeedFrontRight:    0.1,
				SpeedRearLeft:      0.1,
				SpeedRearRight:     0.1,
			}
			updateJson = NewUpdateJson{
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
			vehicleJson := NewVehicle{
				Vin:                "car1",
				Longitude:          previousJsonCar1.Vehicle.Longitude,
				Latitude:           previousJsonCar1.Vehicle.Latitude + previousJsonCar1.Vehicle.SpeedFrontRight*0.0000005,
				DistanceUltrasonic: 100.12,
				DistanceLidar:      100.1,
				SpeedFrontLeft:     previousJsonCar1.Vehicle.SpeedFrontLeft + speedIncrement,
				SpeedFrontRight:    previousJsonCar1.Vehicle.SpeedFrontRight + speedIncrement,
				SpeedRearLeft:      previousJsonCar1.Vehicle.SpeedRearLeft + speedIncrement,
				SpeedRearRight:     previousJsonCar1.Vehicle.SpeedRearRight + speedIncrement,
			}
			updateJson = NewUpdateJson{
				Index:     previousJsonCar1.Index + 1,
				Type:      "update_vehicle",
				Timestamp: "TimestampToReplace",
				Vehicle:   vehicleJson,
			}
			previousJsonCar1 = updateJson

			updateJson.Index = 0

			data = append(data, updateJson)

			// Car 2
			vehicleJson = NewVehicle{
				Vin:                "car2",
				Longitude:          previousJsonCar2.Vehicle.Longitude,
				Latitude:           previousJsonCar2.Vehicle.Latitude + previousJsonCar2.Vehicle.SpeedFrontRight*0.0000005,
				DistanceUltrasonic: previousJsonCar2.Vehicle.DistanceUltrasonic + distanceIncrement,
				DistanceLidar:      previousJsonCar2.Vehicle.DistanceLidar + distanceIncrement,
				SpeedFrontLeft:     previousJsonCar2.Vehicle.SpeedFrontLeft + speedIncrement,
				SpeedFrontRight:    previousJsonCar2.Vehicle.SpeedFrontRight + speedIncrement,
				SpeedRearLeft:      previousJsonCar2.Vehicle.SpeedRearLeft + speedIncrement,
				SpeedRearRight:     previousJsonCar2.Vehicle.SpeedRearRight + speedIncrement,
			}
			updateJson = NewUpdateJson{
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

	// Turning 90 degrees to the West
	for i := 0; i < 9; i++ {
		distanceIncrement := -0.1 + rand.Float64()*(0.1+0.1)

		// Car1
		vehicleJson := NewVehicle{
			Vin:                "car1",
			Longitude:          previousJsonCar1.Vehicle.Longitude - 0.000005,
			Latitude:           previousJsonCar1.Vehicle.Latitude + 0.000005,
			DistanceUltrasonic: 100.12,
			DistanceLidar:      100.1,
			SpeedFrontLeft:     previousJsonCar1.Vehicle.SpeedFrontLeft - 0.000005,
			SpeedFrontRight:    previousJsonCar1.Vehicle.SpeedFrontRight + 0.000005,
			SpeedRearLeft:      previousJsonCar1.Vehicle.SpeedRearLeft,
			SpeedRearRight:     previousJsonCar1.Vehicle.SpeedRearRight,
		}
		updateJson = NewUpdateJson{
			Index:     previousJsonCar1.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar1 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)

		// Car 2
		vehicleJson = NewVehicle{
			Vin:                "car2",
			Longitude:          previousJsonCar2.Vehicle.Longitude - 0.000005,
			Latitude:           previousJsonCar2.Vehicle.Latitude + 0.000005,
			DistanceUltrasonic: previousJsonCar2.Vehicle.DistanceUltrasonic + distanceIncrement,
			DistanceLidar:      previousJsonCar2.Vehicle.DistanceLidar + distanceIncrement,
			SpeedFrontLeft:     previousJsonCar1.Vehicle.SpeedFrontLeft - 0.000005,
			SpeedFrontRight:    previousJsonCar1.Vehicle.SpeedFrontRight + 0.000005,
			SpeedRearLeft:      previousJsonCar1.Vehicle.SpeedRearLeft,
			SpeedRearRight:     previousJsonCar1.Vehicle.SpeedRearRight,
		}
		updateJson = NewUpdateJson{
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
		vehicleJson := NewVehicle{
			Vin:                "car1",
			Longitude:          previousJsonCar1.Vehicle.Longitude - 0.000005,
			Latitude:           previousJsonCar1.Vehicle.Latitude - 0.000005,
			DistanceUltrasonic: 100.12,
			DistanceLidar:      100.1,
			SpeedFrontLeft:     previousJsonCar1.Vehicle.SpeedFrontLeft - 0.000005,
			SpeedFrontRight:    previousJsonCar1.Vehicle.SpeedFrontRight + 0.000005,
			SpeedRearLeft:      previousJsonCar1.Vehicle.SpeedRearLeft,
			SpeedRearRight:     previousJsonCar1.Vehicle.SpeedRearRight,
		}
		updateJson = NewUpdateJson{
			Index:     previousJsonCar1.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar1 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)

		// Car 2
		vehicleJson = NewVehicle{
			Vin:                "car2",
			Longitude:          previousJsonCar2.Vehicle.Longitude - 0.000005,
			Latitude:           previousJsonCar2.Vehicle.Latitude - 0.000005,
			DistanceUltrasonic: previousJsonCar2.Vehicle.DistanceUltrasonic + distanceIncrement,
			DistanceLidar:      previousJsonCar2.Vehicle.DistanceLidar + distanceIncrement,
			SpeedFrontLeft:     previousJsonCar2.Vehicle.SpeedFrontLeft - 0.000005,
			SpeedFrontRight:    previousJsonCar2.Vehicle.SpeedFrontRight + 0.000005,
			SpeedRearLeft:      previousJsonCar2.Vehicle.SpeedRearLeft,
			SpeedRearRight:     previousJsonCar2.Vehicle.SpeedRearRight,
		}
		updateJson = NewUpdateJson{
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
		vehicleJson := NewVehicle{
			Vin:                "car1",
			Longitude:          previousJsonCar1.Vehicle.Longitude,
			Latitude:           previousJsonCar1.Vehicle.Latitude - previousJsonCar1.Vehicle.SpeedFrontLeft*0.00000025,
			DistanceUltrasonic: 100.12,
			DistanceLidar:      100.1,
			SpeedFrontLeft:     previousJsonCar1.Vehicle.SpeedFrontLeft + speedIncrement,
			SpeedFrontRight:    previousJsonCar1.Vehicle.SpeedFrontRight + speedIncrement,
			SpeedRearLeft:      previousJsonCar1.Vehicle.SpeedRearLeft + speedIncrement,
			SpeedRearRight:     previousJsonCar1.Vehicle.SpeedRearRight + speedIncrement,
		}
		updateJson = NewUpdateJson{
			Index:     previousJsonCar1.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar1 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)

		// Car 2
		vehicleJson = NewVehicle{
			Vin:                "car2",
			Longitude:          previousJsonCar2.Vehicle.Longitude,
			Latitude:           previousJsonCar2.Vehicle.Latitude - previousJsonCar2.Vehicle.SpeedFrontLeft*0.00000025,
			DistanceUltrasonic: previousJsonCar2.Vehicle.DistanceUltrasonic + distanceIncrement,
			DistanceLidar:      previousJsonCar2.Vehicle.DistanceLidar + distanceIncrement,
			SpeedFrontLeft:     previousJsonCar2.Vehicle.SpeedFrontLeft + speedIncrement,
			SpeedFrontRight:    previousJsonCar2.Vehicle.SpeedFrontRight + speedIncrement,
			SpeedRearLeft:      previousJsonCar2.Vehicle.SpeedRearLeft + speedIncrement,
			SpeedRearRight:     previousJsonCar2.Vehicle.SpeedRearRight + speedIncrement,
		}
		updateJson = NewUpdateJson{
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
		vehicleJson := NewVehicle{
			Vin:                "car1",
			Longitude:          previousJsonCar1.Vehicle.Longitude + 0.000005,
			Latitude:           previousJsonCar1.Vehicle.Latitude - 0.000005,
			DistanceUltrasonic: 100.12,
			DistanceLidar:      100.1,
			SpeedFrontLeft:     previousJsonCar1.Vehicle.SpeedFrontLeft - 0.000005,
			SpeedFrontRight:    previousJsonCar1.Vehicle.SpeedFrontRight + 0.000005,
			SpeedRearLeft:      previousJsonCar1.Vehicle.SpeedRearLeft,
			SpeedRearRight:     previousJsonCar1.Vehicle.SpeedRearRight,
		}
		updateJson = NewUpdateJson{
			Index:     previousJsonCar1.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar1 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)

		// Car 2
		vehicleJson = NewVehicle{
			Vin:                "car2",
			Longitude:          previousJsonCar2.Vehicle.Longitude + 0.000005,
			Latitude:           previousJsonCar2.Vehicle.Latitude - 0.000005,
			DistanceUltrasonic: previousJsonCar2.Vehicle.DistanceUltrasonic + distanceIncrement,
			DistanceLidar:      previousJsonCar2.Vehicle.DistanceLidar + distanceIncrement,
			SpeedFrontLeft:     previousJsonCar1.Vehicle.SpeedFrontLeft - 0.000005,
			SpeedFrontRight:    previousJsonCar1.Vehicle.SpeedFrontRight + 0.000005,
			SpeedRearLeft:      previousJsonCar1.Vehicle.SpeedRearLeft,
			SpeedRearRight:     previousJsonCar1.Vehicle.SpeedRearRight,
		}
		updateJson = NewUpdateJson{
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
		vehicleJson := NewVehicle{
			Vin:                "car1",
			Longitude:          previousJsonCar1.Vehicle.Longitude + 0.000005,
			Latitude:           previousJsonCar1.Vehicle.Latitude + 0.000005,
			DistanceUltrasonic: 100.12,
			DistanceLidar:      100.1,
			SpeedFrontLeft:     previousJsonCar1.Vehicle.SpeedFrontLeft - 0.000005,
			SpeedFrontRight:    previousJsonCar1.Vehicle.SpeedFrontRight + 0.000005,
			SpeedRearLeft:      previousJsonCar1.Vehicle.SpeedRearLeft,
			SpeedRearRight:     previousJsonCar1.Vehicle.SpeedRearRight,
		}
		updateJson = NewUpdateJson{
			Index:     previousJsonCar1.Index + 1,
			Type:      "update_vehicle",
			Timestamp: "TimestampToReplace",
			Vehicle:   vehicleJson,
		}
		previousJsonCar1 = updateJson

		updateJson.Index = 0

		data = append(data, updateJson)

		// Car 2
		vehicleJson = NewVehicle{
			Vin:                "car2",
			Longitude:          previousJsonCar2.Vehicle.Longitude + 0.000005,
			Latitude:           previousJsonCar2.Vehicle.Latitude + 0.000005,
			DistanceUltrasonic: previousJsonCar2.Vehicle.DistanceUltrasonic + distanceIncrement,
			DistanceLidar:      previousJsonCar2.Vehicle.DistanceLidar + distanceIncrement,
			SpeedFrontLeft:     previousJsonCar2.Vehicle.SpeedFrontLeft - 0.000005,
			SpeedFrontRight:    previousJsonCar2.Vehicle.SpeedFrontRight + 0.000005,
			SpeedRearLeft:      previousJsonCar2.Vehicle.SpeedRearLeft,
			SpeedRearRight:     previousJsonCar2.Vehicle.SpeedRearRight,
		}
		updateJson = NewUpdateJson{
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
