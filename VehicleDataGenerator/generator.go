package VehicleDataGenerator

import (
	"log"
	"os"
)

func GenerateVehicleData() (string, error) {

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

	// Writing data into output file
	err = writeData(logfile, data)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}
