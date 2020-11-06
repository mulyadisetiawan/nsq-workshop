package nsq

import (
	"github.com/bitly/go-nsq"
	"github.com/sharring_session/nsq/pkg/mq"
)

type subscriber struct {
	topic   string
	channel string
	handler nsq.Handler
}

// MQ struct
type MQ struct {
	options   *Options
	consumers []subscriber
}

// Options struct
type Options struct {
	ListenAddress  string
	PublishAddress string
	Prefix         string
}

func New(o *Options) *MQ {
	m := &MQ{
		options: o,
	}
	return m
}

func (m *MQ) RegisterSubcribers(subs mq.Subscribers) error {
	for topicName, topic := range subs {
		for channelName, channel := range topic {
			m.consumers = append(m.consumers, subscriber{
				topic:   topicName,
				channel: channelName,
				handler: channel.Handler,
			})
		}
	}
	return nil
}

func (m *MQ) Run() error {
	if m.options.PublishAddress == "" {
		m.options.PublishAddress = m.options.ListenAddress
	}
	for _, consumer := range m.consumers {
		q, err := nsq.NewConsumer(consumer.topic, consumer.channel, nsq.NewConfig())
		if err != nil {
			return err
		}

		q.AddHandler(consumer.handler)

		err = q.ConnectToNSQD(m.options.ListenAddress)
		if err != nil {
			return err
		}
	}
	return nil
}
