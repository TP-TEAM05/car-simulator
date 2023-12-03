package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type JsonLogLine struct {
	Level          string `json:"level"`
	ReceivingPort  int    `json:"receivingPort"`
	ConnectionType string `json:"connectionType"`
	SourceIP       string `json:"sourceIP"`
	SourcePort     int    `json:"sourcePort"`
	Time           string `json:"time"`
	Message        string `json:"message"`
}

// Index correction variables
var vehicleOneIndex = 1
var vehicleTwoIndex = 1
var isFirst = true

func StartProcessingJSON(dumpFilepath string, startTimeOffset float32, connectionsManager *ConnectionsManager) {

	// Open dump file
	dumpJsonFile, err := os.Open(dumpFilepath)
	if err != nil {
		fmt.Println("Error occurred while opening the file:", err)
		return
	}
	defer dumpJsonFile.Close()

	var startTime = time.Now().UTC()
	var dumpStartTime *time.Time
	var readDumpFromTime time.Time

	scanner := bufio.NewScanner(dumpJsonFile)

	var correctedTimestamp string

	// Read dump line by line
	for scanner.Scan() {
		lineBytes := scanner.Bytes()
		var jsonLine JsonLogLine
		err := json.Unmarshal(lineBytes, &jsonLine)
		if err != nil {
			fmt.Printf("Error unmarshalling json dump file line: %v\n", err)
			continue
		}

		/*
















		 */
		if dumpFilepath == os.Getenv("CONNECT_FILE_PATH") {
			err = setTimestampEpoch(jsonLine.Time)
			if err != nil {
				fmt.Println("Error occurred while setting epoch:", err)
				return
			}
		} else {
			correctedTimestamp = getCorrectedTimestamp()
			jsonLine.Time = correctedTimestamp
		}
		/*



















		 */

		// Get the time of the first line of the dump
		if dumpStartTime == nil {
			parsedTime, err := time.Parse(TimestampFormat, jsonLine.Time)
			if err != nil {
				fmt.Printf("Failed to parse time of first dump line: %v\n", err)
				continue
			}
			dumpStartTime = &parsedTime
			readDumpFromTime = dumpStartTime.Add(time.Duration(float32(time.Second) * startTimeOffset))
		}

		// Timestamp of current log line
		currentLineTime, err := time.Parse(TimestampFormat, jsonLine.Time)
		if err != nil {
			fmt.Printf("Failed to parse time of dump line: %v\n", err)
			continue
		}

		// Start sending dump only from specified time offset
		if currentLineTime.Before(readDumpFromTime) {
			continue
		}

		// Get Base Datagram
		var baseDatagram BaseDatagram
		err = json.Unmarshal([]byte(jsonLine.Message), &baseDatagram)
		if err != nil {
			fmt.Printf("Error unmarshalling json message to datagram: %v\n", err)
			continue
		}

		/*

















		 */

		// Index and timestamp correction
		if baseDatagram.Type == "update_vehicle" {
			if isFirst {
				jsonLine.Message = strings.Replace(jsonLine.Message, "0", strconv.Itoa(vehicleOneIndex), 1)
				vehicleOneIndex += 1
				isFirst = false
			} else {
				jsonLine.Message = strings.Replace(jsonLine.Message, "0", strconv.Itoa(vehicleTwoIndex), 1)
				vehicleTwoIndex += 1
				isFirst = true
			}

			jsonLine.Message = strings.Replace(jsonLine.Message, "TimestampToReplace", correctedTimestamp, 1)
		}

		/*



















		 */

		datagramToSend := []byte(jsonLine.Message)
		var vin string

		// Get Specific Datagram
		switch baseDatagram.Type {
		case "update_vehicle":
			var updateVehicleDatagram UpdateVehicleDatagram
			err = json.Unmarshal(datagramToSend, &updateVehicleDatagram)
			if err != nil {
				fmt.Printf("Error unmarshalling json message to UpdateVehicleDatagram datagram: %v\n", err)
				continue
			}
			vin = updateVehicleDatagram.Vehicle.Vin

		case "connect_vehicle":
			var connectVehicleDatagram ConnectVehicleDatagram
			err = json.Unmarshal(datagramToSend, &connectVehicleDatagram)
			if err != nil {
				fmt.Printf("Error unmarshalling json message to ConnectVehicleDatagram datagram: %v\n", err)
				continue
			}
			vin = connectVehicleDatagram.Vin
		default:
			continue
		}

		// Wait until time is right
		var timestamp = startTime.Add(currentLineTime.Sub(readDumpFromTime))
		time.Sleep(time.Until(timestamp))

		connection := connectionsManager.GetOrCreateConnection(vin)
		connection.WriteDatagram(datagramToSend)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading json dump file: %v\n", err)
	}
}
