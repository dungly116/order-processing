package kafkaconfig

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Brokers []string `json:"brokers"`
	Topic   string   `json:"topic"`
}

func LoadKafkaConfig(filename string) (Config, error) {
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	var kafkaConfig Config
	err = json.Unmarshal(configFile, &kafkaConfig)
	if err != nil {
		return Config{}, err
	}

	return kafkaConfig, nil
}
