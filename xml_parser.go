package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"time"
)

// Structure
//	<timestep time="5.00">
//		<vehicle id="veh0" x="17.160442" y="48.362205" angle="99.31" type="veh_passenger" speed="9.30" pos="32.15" lane="-1043858896_0" slope="0.00"/>
//  	<vehicle id="veh0.1" x="17.160170" y="48.362237" angle="98.36" type="veh_passenger" speed="4.01" pos="11.68" lane="-1043858896_0" slope="0.00"/>
// 	</timestep>

// XML structs

type TimestepXML struct {
	Time     float32      `xml:"time,attr"`
	Vehicles []VehicleXML `xml:"vehicle"`
}

type VehicleXML struct {
	Id     string  `xml:"id,attr"`
	X      float32 `xml:"x,attr"`
	Y      float32 `xml:"y,attr"`
	Angle  float32 `xml:"angle,attr"`
	Type   string  `xml:"type,attr"`
	Speed  float32 `xml:"speed,attr"`
	LaneId string  `xml:"lane,attr"`
}

func vehicleXMLToVehicleJSON(vehicleXML *VehicleXML) UpdateVehicleVehicle {
	return UpdateVehicleVehicle{
		Vin:   vehicleXML.Id,
		Speed: vehicleXML.Speed,
	}
}

func StartProcessingXML(dumpFilepath string, startTimeOffset float32, connectionsManager *ConnectionsManager) {

	// Open dump file
	dumpXmlFile, err := os.Open(dumpFilepath)
	if err != nil {
		fmt.Println("Error occurred while opening the file:", err)
		return
	}
	defer dumpXmlFile.Close()

	var xmlDecoder = xml.NewDecoder(dumpXmlFile)
	var startTime = time.Now().UTC()

	for {
		token, _ := xmlDecoder.Token()
		if token == nil {
			break
		}

		switch tokenType := token.(type) {
		case xml.StartElement:
			if tokenType.Name.Local == "timestep" {
				var timestep TimestepXML
				_ = xmlDecoder.DecodeElement(&timestep, &tokenType)

				timestepAdjustedTime := timestep.Time - startTimeOffset
				if timestepAdjustedTime < 0 {
					continue
				}

				// Send just at time
				var timestamp = startTime.Add(time.Duration(timestepAdjustedTime * float32(time.Second)))
				time.Sleep(time.Until(timestamp))

				var currentVehicleIds []string

				// NOTE: We don't send here Connect Vehicle datagram, it is not required, but in non-simulator environment
				// it may be beneficial (more info in the API Documentation)
				for _, vehicle := range timestep.Vehicles {
					vehicleJson := vehicleXMLToVehicleJSON(&vehicle)
					currentVehicleIds = append(currentVehicleIds, vehicleJson.Vin)

					connection := connectionsManager.GetOrCreateConnection(vehicleJson.Vin)

					datagram, _ := json.Marshal(UpdateVehicleDatagram{
						BaseDatagram: BaseDatagram{
							Index:     connection.NextSendIndex,
							Timestamp: timestamp.Format(TimestampFormat),
							Type:      "update_vehicle",
						},
						Vehicle: vehicleJson,
					})

					connection.WriteDatagram(datagram)
				}

				// Clean unused connections
				connectionsManager.DeleteAllConnectionsExcept(currentVehicleIds)
			}
		}
	}
}
