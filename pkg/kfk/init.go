package kfk

import (
	"SearchEngine/config"

	log "SearchEngine/pkg/logger"

	"github.com/IBM/sarama"
)

var GobalKafka sarama.Client

func InitKafka() {
	con := sarama.NewConfig()
	con.Producer.Return.Successes = true
	kafkaClient, err := sarama.NewClient(config.Conf.Kafka.Address, con)
	if err != nil {
		log.LogrusObj.Errorln(err)
		return
	}
	GobalKafka = kafkaClient
}
