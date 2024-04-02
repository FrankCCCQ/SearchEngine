package kfk

import (
	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

// KafkaProducer 发送单条消息
func KafkaProducer(topic string, msg []byte) (err error) {
	producer, err := sarama.NewSyncProducerFromClient(GobalKafka)
	if err != nil {
		return errors.Wrap(err, "failed to create kafka producer")
	}
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	_, _, err = producer.SendMessage(message)
	if err != nil {
		return errors.Wrap(err, "failed to send message")
	}
	return
}

func KafkaProducers(messages []*sarama.ProducerMessage) (err error) {
	producer, err := sarama.NewSyncProducerFromClient(GobalKafka)
	if err != nil {
		return errors.Wrap(err, "failed to create kafka producer")
	}
	err = producer.SendMessages(messages)
	if err != nil {
		return errors.Wrap(err, "failed to send messages")
	}
	return
}
