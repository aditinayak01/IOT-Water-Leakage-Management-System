package logger

import (
	"log"
	"os"
	"github.com/ambatshri/mqtt_goclient/config"
	"github.com/ambatshri/mqtt_goclient/utils"
	"fmt"
)
var Levels = []string{"DEBUG","INFO","WARN","ERROR"}
var logLevelInt int
func init() {
	logLevel := config.Config.SealClient.Loglevel
	logLevelInt = utils.ArrayIndex(Levels, logLevel)
	f, err := os.OpenFile(config.Config.SealClient.LogDirectory+"/system.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
        	os.Exit(1)
	} else {
        	log.SetOutput(f)
	}
}

func Write(level string, message interface{}) {
	levelInt := utils.ArrayIndex(Levels,level)
	if levelInt >= logLevelInt && logLevelInt != -1 {
		put(level, message)
	}	
}

func put(level string, msg interface{}) {
	switch dataType := msg.(type) {
		case string:
			message := " [ " + level + " ] " + msg.(string)
			log.Printf(message)
			_ = dataType
	}
}
