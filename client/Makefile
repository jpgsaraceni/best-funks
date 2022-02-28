.PHONY: producer consumer

producer:
	go build -o out/producer producer/producer.go
	./out/producer getting-started.properties

consumer: 
	go build -o out/consumer consumer/consumer.go
	./out/consumer getting-started.properties