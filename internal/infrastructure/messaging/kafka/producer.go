package kafka

import (
	"context"
	"encoding/json"

	"github.com/Gadzhimurad17/link-checker/internal/domain/models"
	"github.com/Gadzhimurad17/link-checker/internal/infrastructure/dto"
	"github.com/IBM/sarama"
)

type Producer struct {
	producer sarama.AsyncProducer
	topics   models.Topic
}

func NewProducer(addrs []string, conf *sarama.Config) *Producer {
	asyncProducer, _ := sarama.NewAsyncProducer(addrs, conf)

	return &Producer{
		producer: asyncProducer,
	}
}

func (p *Producer) newLinkMessage(event *dto.LinkLog) (*sarama.ProducerMessage, error) {
	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic: string(models.LinkTopic),
		Value: sarama.ByteEncoder(data),
	}, nil
}

func (p *Producer) PublishWithContext(ctx context.Context, event *dto.LinkLog){
	//TODO publish with context
}

func (p *Producer) Close() error {
	return p.producer.Close()
}
