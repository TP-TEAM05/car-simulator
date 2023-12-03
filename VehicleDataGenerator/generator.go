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
	data, err := generateData()
	if err != nil {
		return "", err
	}

	// Writing data into connect output file
	err = writeData(connectLogfile, data, true)
	if err != nil {
		return "", err
	}

	// Writing data into output file
	err = writeData(logfile, data, false)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
