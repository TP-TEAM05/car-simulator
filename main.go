package main

import (
	"car-simulator/VehicleDataGenerator"
	"flag"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
)

func loadServerAddress() *net.UDPAddr {
	// Get IP Address of Integration Module
	ips, err := net.LookupIP(os.Getenv("INTEGRATION_MODULE_HOST"))
	if err != nil || len(ips) == 0 {
		fmt.Printf("Could not resolve or find hostname %v", err)
		os.Exit(1)
	}

	// Get Port of Integration Module
	port, err := strconv.Atoi(os.Getenv("INTEGRATION_MODULE_PORT"))
	if err != nil {
		fmt.Printf("Failed to parse port %v", err)
	}
	return &net.UDPAddr{Port: port, IP: ips[0]}
}

func main() {

	// Load the cmd arguments
	flag.Parse()

	dumpFilepath := os.Getenv("VEHICLE_DUMP_FILEPATH")
	startTimeOffset := GetFloatFromEnv("VEHICLE_DUMP_START_TIME_OFFSET", 0)

	// Generating data
	dumpFilepath, err := VehicleDataGenerator.GenerateVehicleData()
	if err != nil {
		log.Println(err)
		return
	}

	// Get IP and Port of the Integration Module
	serverAddress := loadServerAddress()
	connectionsManager := NewConnectionsManager(serverAddress)

	switch filepath.Ext(dumpFilepath) {
	case ".xml":
		StartProcessingXML(dumpFilepath, startTimeOffset, connectionsManager)
	case ".log":
		for true {
			StartProcessingJSON(dumpFilepath, startTimeOffset, connectionsManager)
		}
	default:
		panic(fmt.Sprintf("unknown dump file format with extension %s", filepath.Ext(dumpFilepath)))
	}
}
