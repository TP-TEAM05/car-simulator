package VehicleDataGenerator

import (
	"log"
	"os"
)

func GenerateVehicleData() (string, error) {

	// Creation of connection logfile
	connectOutputPath := os.Getenv("CONNECT_FILE_PATH")
	connectLogfile, err := os.Create(connectOutputPath)
	if err != nil {
		return "", err
	}
	defer func(connectLogfile *os.File) {
		err := connectLogfile.Close()
		if err != nil {
			log.Println(err)
		}
	}(connectLogfile)

	// Creation of logfile
	outputPath := os.Getenv("OUTPUT_PATH")
	logfile, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer func(logfile *os.File) {
		err := logfile.Close()
		if err != nil {
			log.Println(err)
		}
	}(logfile)

	// Generating data
	wantNewData, err := GetFloatFromEnv("WANT_NEW_DATA")
	if err != nil {
		return "", err
	}

	var data []interface{}
	if wantNewData == 1 {
		data, err = generateNewData()
		if err != nil {
			return "", err
		}
	} else {
		data, err = generateData()
		if err != nil {
			return "", err
		}
	}

	if wantNewData == 1 {
		// Writing data into connect output file
		err = writeData(connectLogfile, data, true, true)
		if err != nil {
			return "", err
		}

		// Writing data into output file
		err = writeData(logfile, data, false, true)
		if err != nil {
			return "", err
		}
	} else {
		// Writing data into connect output file
		err = writeData(connectLogfile, data, true, false)
		if err != nil {
			return "", err
		}

		// Writing data into output file
		err = writeData(logfile, data, false, false)
		if err != nil {
			return "", err
		}
	}

	return outputPath, nil
}
