package util

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ReadConfig() kafka.ConfigMap {
	// check if config file was passed as argument in command line
	if len(os.Args) != 2 {
		log.Fatalf("missing .properties config file argument. Usage: %s <path to config file>\n", os.Args[0])
	}
	// get config file name from command line argument
	configFile := os.Args[1]

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
