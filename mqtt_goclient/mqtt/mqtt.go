package mqtt
import (
  "fmt"
  MQTT "github.com/eclipse/paho.mqtt.golang"
  "os"
  "time"
  "github.com/ambatshri/mqtt_goclient/config"
)

//define a function for the default message handler
var MessageHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fz,err := os.OpenFile("testlog", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
	}
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
	defer fz.Close()
}

func getClientOptions() *MQTT.ClientOptions {
	opts := MQTT.NewClientOptions().AddBroker(config.Config.MsgQueue.Protocol + "://" + config.Config.MsgQueue.Host + ":" + config.Config.MsgQueue.Port)
	opts.SetClientID(config.Config.MsgQueue.Clientid)
	opts.SetDefaultPublishHandler(MessageHandler)
	return opts
}

func NewClient() MQTT.Client {
	return MQTT.NewClient(getClientOptions())
}

func Connect(msgQueue MQTT.Client) error {
	if token := msgQueue.Connect(); token.Wait() && token.Error() != nil {
		// don't panic return error
		return token.Error()
	}
	return nil
}

func Process(msgQueue MQTT.Client) error {
	time.Sleep(3*time.Second)
	if token := msgQueue.Subscribe(config.Config.MsgQueue.Topic, 0, nil); token.Wait() && token.Error() != nil {
  		return token.Error()
	}
	return nil
}
