package writer

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	writer *kafka.Writer
	topic  string
	key    string
}

func (k *Kafka) Write(b []byte) (int, error) {
	var key []byte
	copy(key, k.key)
	err := k.writer.WriteMessages(context.Background(), kafka.Message{
		Topic: k.topic,
		Key:   key,
		Value: b,
	})
	if err != nil {
		return 0, err
	}
	return len(b), nil
}

func NewKafkaWriter(w *kafka.Writer, topic, key string) *Kafka {
	return &Kafka{
		writer: w,
		topic:  topic,
		key:    key,
	}
}
