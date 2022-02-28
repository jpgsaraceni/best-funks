package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/jpgsaraceni/best-funks/client/util"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s ~/best-funks/client/getting-started.properties\n", os.Args[0])
	}
	configFile := os.Args[1]
	conf := util.ReadConfig(configFile)

	topic := "purchases"
	p, err := kafka.NewProducer(&conf)

	if err != nil {
		log.Fatalf("failed to create producer: %s", err)
	}

	// Go-routine to handle message delivery reports and
	// possibly other event types (errors, stats, etc)
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Fatalf("failed to deliver message: %v\n", ev.TopicPartition)
				}
				log.Printf(
					"Produced event to topic %s: key = %-10s value = %s\n",
					*ev.TopicPartition.Topic,
					string(ev.Key),
					string(ev.Value),
				)
			}
		}
	}()

	users := [...]string{"eabara", "jsmith", "sgarcia", "jbernard", "htanaka", "awalther"}
	items := [...]string{"book", "alarm clock", "t-shirts", "gift card", "batteries"}

	// produce random messages using of a user(key) and item(value)
	for n := 0; n < 10; n++ {
		rand.Seed(time.Now().UnixNano())
		key := users[rand.Intn(len(users))]
		data := items[rand.Intn(len(items))]
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(key),
			Value:          []byte(data),
		}, nil)
	}

	// Wait for all messages to be delivered
	p.Flush(15 * 1000)
	p.Close()
}
