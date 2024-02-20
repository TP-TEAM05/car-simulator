package VehicleDataGenerator

import (
	"encoding/json"
	"os"
	"strings"
)

func writeData(logfile *os.File, data []interface{}, toConnect bool, newData bool) error {

	if toConnect {
		for i := 0; i < 2; i++ {
			model := data[i]
			jsonBytes, err := json.Marshal(model)
			if err != nil {
				return err
			}

			var timestamp string

			if _, ok := model.(ConnectJson); ok {
				timestamp = model.(ConnectJson).Timestamp
			}

			if newData {
				if _, ok := model.(NewUpdateJson); ok {
					timestamp = model.(NewUpdateJson).Timestamp
				}
			} else {
				if _, ok := model.(UpdateJson); ok {
					timestamp = model.(UpdateJson).Timestamp
				}
			}

			jsonString := string(jsonBytes)
			jsonString = strings.ReplaceAll(jsonString, "\"", "\\\"")
			jsonString = "{\"time\":\"" + timestamp + "\",\"message\":\"" + jsonString + "\"}"

			_, err = logfile.WriteString(jsonString + "\n")
			if err != nil {
				return err
			}
		}
	} else {
		for i, model := range data {
			if i > 1 {
				jsonBytes, err := json.Marshal(model)
				if err != nil {
					return err
				}

				var timestamp string

				if newData {
					if _, ok := model.(NewUpdateJson); ok {
						timestamp = model.(NewUpdateJson).Timestamp
					}
				} else {
					if _, ok := model.(UpdateJson); ok {
						timestamp = model.(UpdateJson).Timestamp
					}
				}

				jsonString := string(jsonBytes)
				jsonString = strings.ReplaceAll(jsonString, "\"", "\\\"")
				jsonString = "{\"time\":\"" + timestamp + "\",\"message\":\"" + jsonString + "\"}"

				_, err = logfile.WriteString(jsonString + "\n")
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
