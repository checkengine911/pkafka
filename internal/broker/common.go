package broker

import (
	"kafkago/internal/broker/request"
	"kafkago/pkg/kafka"
)

type Brokers struct {
	Producer *request.Producer
}
//for init broker
func InitBrokers(dialer *kafka.Dialer) (*Brokers, error) {
	return &Brokers{
		Producer: request.NewProducer(dialer),
	}, nil
}
