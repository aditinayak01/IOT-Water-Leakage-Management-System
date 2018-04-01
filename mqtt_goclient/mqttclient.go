package main

import (
  "fmt"
  //import the Paho Go MQTT library
  MQTT "github.com/eclipse/paho.mqtt.golang"
  "os"
  "time"
  "log"
)

//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
  fmt.Printf("TOPIC: %s\n", msg.Topic())
  fmt.Printf("MSG: %s\n", msg.Payload())
  fz,err := os.OpenFile("testlog", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
    fmt.Printf("Error opening file: %v", err)
  }
  defer fz.Close()
  log.SetOutput(fz)
  log.Println("Msg:",string(msg.Payload()))
}

func main() {
  //create a ClientOptions struct setting the broker address, clientid, turn
  //off trace output and set the default message handler
  opts := MQTT.NewClientOptions().AddBroker("tcp://172.17.0.3:1883")
  opts.SetClientID("go-simple")
  opts.SetDefaultPublishHandler(f)

  //create and start a client using the above ClientOptions
  c := MQTT.NewClient(opts)
  if token := c.Connect(); token.Wait() && token.Error() != nil {
    panic(token.Error())
  }

  //subscribe to the topic /go-mqtt/sample and request messages to be delivered
  //at a maximum qos of zero, wait for the receipt to confirm the subscription
for{
time.Sleep(3*time.Second)  
if token := c.Subscribe("outTopic", 0, nil); token.Wait() && token.Error() != nil {
    fmt.Println(token.Error())
    os.Exit(1)
  }
}
  //Publish 5 messages to /go-mqtt/sample at qos 1 and wait for the receipt
  //from the server after sending each message

  time.Sleep(10 * time.Second)

  //unsubscribe from /go-mqtt/sample
  if token := c.Unsubscribe("outTopic"); token.Wait() && token.Error() != nil {
    fmt.Println(token.Error())
    os.Exit(1)
  }

  c.Disconnect(250)
}
