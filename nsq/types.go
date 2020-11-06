package nsqmodule

import nsq "github.com/nsqio/go-nsq"

// Producer
type (
	ProducerConfig struct {
		NsqdAddress string
	}

	Producer struct {
		prod *nsq.Producer
	}
)

// Consumer
type (
	ConsumerConfig struct {
		Topic         string
		Channel       string
		LookupAddress string
		MaxAttempts   uint16
		MaxInFlight   int
		Handler       nsq.HandlerFunc
	}

	Consumer struct {
		cons          *nsq.Consumer
		lookupAddress string
		handler       nsq.HandlerFunc
	}
)
