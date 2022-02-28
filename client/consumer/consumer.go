package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jpgsaraceni/best-funks/client/util"
)

func main() {
	conf := util.ReadConfig()

	conf["group.id"] = "best-funks"
	conf["auto.offset.reset"] = "earliest"

	c, err := kafka.NewConsumer(&conf)

	if err != nil {
		log.Fatalf("failed to create consumer: %s", err)
	}

	topic := "purchases"
	if err = c.SubscribeTopics([]string{topic}, nil); err != nil {
		log.Fatalf("failed to subscribe to topic 'purchases': %s", err)
	}
	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			log.Printf(
				"Consumed event from topic %s key = %-10s value = %s\n",
				*ev.TopicPartition.Topic,
				string(ev.Key),
				string(ev.Value),
			)
		}
	}

	c.Close()
}
