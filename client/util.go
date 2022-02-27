package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ReadConfig(configFile string) kafka.ConfigMap {
	// map that will receive config key-value pairs
	m := make(map[string]kafka.ConfigValue)

	// open config file
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// scan config file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// ignore comments and blank lines
		if !strings.HasPrefix(line, "#") && len(line) != 0 {
			// map key-values of config file
			kv := strings.Split(line, "=")
			paramater := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[paramater] = value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	return m
}
