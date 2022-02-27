# Step by step

## 1. Run a Kafka Cluster

You can use a Confluent Cloud cluster, or a docker image (like I did). Just copy the `docker-compose.yml` file and run `docker-compose up` in your terminal.

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

## 4. Create `utils.go`

This file contains helper functions for client configurations. It is (almost) a copy from the tutorial.

## 5. Create `producer.go`

This creates the Kafka procucer.
