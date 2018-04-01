package config
import (
	"github.com/spf13/viper"
	"github.com/ambatshri/mqtt_goclient/constant"
	"fmt"
)
// create a struct config similar to config.toml
// use viper to get store config toml values in this struct
// this config variable will be global

type sealclient struct {
	Loglevel string
	LogDirectory string
}

type msgqueue struct {
	Name string
	Host string
	Protocol string
	Port string
	Topic string
	Clientid string
} 

type config struct {
	SealClient sealclient
	MsgQueue msgqueue	
}

var Config config


func init() {
	err := getLocal()
	if err != nil {
		fmt.Println(err.Error())
	}
}


func getLocal() error {

	viper.SetConfigName("config") // path to look for the config file in
	viper.AddConfigPath(constant.ConfigPath)
	viper.SetConfigType("toml")

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}
	
	viper.Unmarshal(&Config)
	return nil
}
