# Apache Kafka

My study repo for Apache Kafka. Based on [this tutorial](https://kafka.apache.org/intro).

## Contents

* [Overview](#overview)
* [Key Terms](#key-terms)
  * [Event](#event)
  * [Topic](#topic)
  * [Producer](#producer)
  * [Consumer](#consumer)
  * [Partition](#partition)
* [Getting Started](#getting-started)
  * [Install](#install)
  * [Run](#run)
  * [Clear Data](#clear-data)
* [Kafka Clients](#kafka-clients)

## Overview

Kafka is a distributed event streaming service. Events (also called "records" or "messages" - things that happened) are stored in "topics". Topics are persisted to disk (for a definite or indefinite period of time). Instead of storing a state of an object in a database, it stores events (and their states) in logs (called "topics" in Kafka).

Kafka is especially useful in microservices, for communication between services (producing and consuming events from one another). It also allows for real-time analytics of events, in a more straightforward manner.

Kafka Connect gets data from a database and sets it to a topic.

Kafka Streams is a Java API that does services like aggregation, grouping, enrichment (joins) on Kafka topics.

## Key Terms

### Event

Things that happened. Also refered to as messages or records, they are represented by a key, a value, a timestamp and optional metadata.

### Topic

A log of events.

### Producer

Client applications that write (publish) to a topic.

### Consumer

Client applications that read (subscribe) from a topic.

### Partition

Parts of a topic apread over buckets on Kafka brokers. Events with the same event key are always stored in the same partition. Consumers of a given topic-partition always read events in the order they are written.

## Getting Started

Based on the [step-by-step guide](https://kafka.apache.org/quickstart) on the official website.

### Install

1. Download the suggested version from the link above.

2. Extract:

    ```shell
    tar -xzf kafka_2.13-3.1.0.tgz
    cd kafka_2.13-3.1.0
    ```

3. Install Java (requires 8+):

    ```shell
    sudo apt update
    sudo apt install default-jre
    java -version
    ```

### Run

1. Start Kafka environment:

    Start the ZooKeeper server:

    ```shell
    bin/zookeeper-server-start.sh config/zookeeper.properties
    ```

    And in another terminal instance, start the Kafka broker:

    ```shell
    bin/kafka-server-start.sh config/server.properties
    ```

2. Create a topic (in another terminal instance):

    ```shell
    bin/kafka-topics.sh --create --topic quickstart-events --bootstrap-server localhost:9092
    ```

    Run `kafka-topics.sh` to display usage information.

3. Write events into the topic:

    Run the console producer client:

    ```shell
    bin/kafka-console-producer.sh --topic quickstart-events --bootstrap-server localhost:9092
    ```

    Enter your events:

    ```console
    >An event
    >Another event
    ```

    enter ctrl+c to exit

4. Read the events:

    Run the console consumer client:

    ```shell
    bin/kafka-console-consumer.sh --topic quickstart-events --from-beginning --bootstrap-server localhost:9092
    ```

### Clear data

To remove created topics and events:

```shell
rm -rf /tmp/kafka-logs /tmp/zookeeper
```

## Kafka Clients

There are a variety of [clients available](https://cwiki.apache.org/confluence/display/KAFKA/Clients) for using Kafka from inside an application. In the [client directory of this repo](https://github.com/jpgsaraceni/best-funks/client) I will build an implementation of [Confluent's Go client](https://docs.confluent.io/clients-confluent-kafka-go/current/overview.html), following their [tutorial](https://developer.confluent.io/get-started/go?_ga=2.27437308.1727033679.1645980141-1350969239.1645980141).
