package main

import (
  "github.com/ambatshri/mqtt_goclient/mqtt"
  "fmt"
  "github.com/ambatshri/mqtt_goclient/logger"
)

func main() {
	// create a new client
	mqttClient := mqtt.NewClient()
	logger.Write("ERROR","Seal")	
	// connect
	err := mqtt.Connect(mqttClient)
	if err != nil {
		// log error and return
		fmt.Println(err)
		return
	}
	
	for {
		// process
		mqtt.Process(mqttClient)
	}
}
