package config

type BrokerConfig struct {
	KafkaBrokersPort   int    `env:"KAFKA_BROKERS"`
	KafkaBrokersHost   string `env:"KAFKA_BROKERS_HOST"`
	KafkaClientId      string `env:"KAFKA_CLIENT_ID"`
	KaflaLogTopic      string `env:"KAFKA_LOG_TOPIC"`
	KafkaConsumerGroup string `env:"KAFKA_CONSUMER_GROUP"`
	LogFallbackPath    string `env:"LOG_FALLBACK_PATH"`
}
