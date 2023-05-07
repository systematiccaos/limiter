package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/systematiccaos/going-forward/mqtt"
	"github.com/systematiccaos/going-forward/util"
)

var last_total_power float64 = 0.0
var last_solar_power float64 = 0.0

func main() {
	util.SetupLogs()
	client := mqtt.Client{}
	client_id := os.Getenv("MQTT_CLIENT_ID")
	broker := os.Getenv("MQTT_URL")
	user := os.Getenv("MQTT_USER")
	password := os.Getenv("MQTT_PASSWORD")
	publish_topic := os.Getenv("MQTT_PUB_TOPIC")
	logrus.Println(broker)
	if err := client.Connect(broker, client_id, user, password, true); err != nil {
		logrus.Fatalf("could not connect to mqtt! %s", err)
	}
	power_chan := make(chan mqtt.MQTTSubscriptionMessage)

	if err := client.Subscribe("#", power_chan); err != nil {
		logrus.Fatalln("could not subscribe to power_chan - check MQTT_TOPIC_SOLAR_POWER")
	}
	go rxPower(power_chan)
	for {
		limit := 700.0
		if last_total_power < 10 && last_total_power != 0.0 {
			limit = last_total_power + last_solar_power - 10
		}
		logrus.Printf("limit: %f", limit)
		logrus.Printf("last_total_power: %f", last_total_power)
		logrus.Printf("last_solar_power: %f", last_solar_power)
		tk := client.Publish(publish_topic, fmt.Sprintf("%f", limit))
		if tk.Error() != nil {
			logrus.Errorln(tk.Error())
		}
		time.Sleep(20 * time.Second)
	}
}

func rxPower(power_chan chan mqtt.MQTTSubscriptionMessage) {
	topic_total_power := os.Getenv("MQTT_TOPIC_TOTAL_POWER")
	topic_solar_power := os.Getenv("MQTT_TOPIC_SOLAR_POWER")

	for {
		msg, more := <-power_chan
		if !more {
			logrus.Fatalln("channel was closed")
		}
		if msg.Message.Topic() == topic_total_power {
			var err error
			last_total_power, err = strconv.ParseFloat(string(msg.Message.Payload()), 64)
			if err != nil {
				logrus.Errorln(err)
			}
		}
		if msg.Message.Topic() == topic_solar_power {
			var err error
			last_solar_power, err = strconv.ParseFloat(string(msg.Message.Payload()), 64)
			if err != nil {
				logrus.Errorln(err)
			}
		}
		msg.Message.Ack()
	}
}
