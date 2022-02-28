# Step by step

## 1. Run a Kafka Cluster

You can use a Confluent Cloud cluster, or a docker image (like I did). Just copy the `docker-compose.yml` file and run `docker-compose up` in your terminal (you will need [Docker Compose](https://docs.docker.com/compose/install/) installed).

## 2. Add Configuration File

Copy the `getting-started.properties` file.

## 3. Create a Topic

Use the kafka broker to create a topic called `purchases`:

```shell
docker-compose exec broker \
  kafka-topics --create \
    --topic purchases \
    --bootstrap-server localhost:9092 \
    --replication-factor 1 \
    --partitions 1
```

## 4. Create `utils/utils.go`

This file contains helper functions for client configurations. It is (almost) a copy from the tutorial.

## 5. Create `producer/producer.go`

This creates the Kafka procucer. You can run the producer using `make producer` (if you've cloned this repo or copied the Makefile).

## 6. Create `consumer/consumer.go`

This creates the Kafka consumer. With Makefile, run `make consumer`.
