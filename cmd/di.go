package cmd

import "github.com/ankitchauhan123/kafka-go/cmd/lib"

type DIContainer struct {
	appName string
	env     lib.Env

	producerSimple func() (*kafkautils.SimpleProducer, error)
	producerBatch  func() (*kafkautils.SimpleProducer, error)
}
