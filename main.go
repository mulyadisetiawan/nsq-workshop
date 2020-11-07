package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bitly/go-nsq"
	"github.com/sharring_session/nsq/config"
	benefithttp "github.com/sharring_session/nsq/handler/http/benefit"
	ovonsq "github.com/sharring_session/nsq/handler/nsq/ovo"
	mqlib "github.com/sharring_session/nsq/pkg/mq"
	nsqlib "github.com/sharring_session/nsq/pkg/mq/nsq"
	benefitrepo "github.com/sharring_session/nsq/repository/benefit"
	ovorepo "github.com/sharring_session/nsq/repository/ovo"
	benefituc "github.com/sharring_session/nsq/usecase/benefit"
	ovouc "github.com/sharring_session/nsq/usecase/ovo"
)

var (
	benefitRepo *benefitrepo.Repository
	ovoRepo     *ovorepo.Repository

	benefitUsecase *benefituc.Usecase
	ovoUsecase     *ovouc.Usecase
)

func main() {
	fmt.Println("RUNNING")

	producer, err := nsq.NewProducer(config.NSQD, nsq.NewConfig())
	if err != nil {
		log.Fatal(err)
	}

	benefitRepo = benefitrepo.NewRepository(
		producer,
	)
	ovoRepo = ovorepo.NewRepository()

	benefitUsecase = benefituc.NewUsecase(
		benefitRepo,
	)
	ovoUsecase = ovouc.NewUsecase(
		ovoRepo,
	)

	benefithttp.NewHandler(benefitUsecase)
	ovoNSQ := ovonsq.NewHandler(ovoUsecase)

	mq := nsqlib.New(&nsqlib.Options{
		ListenAddress: config.NSQD,
	})

	mq.RegisterSubcribers(mqlib.Subscribers{
		config.NSQTopic: {
			config.NSQChannel: mqlib.SubscriberOptions{
				Handler: mqlib.NewNSQLHandler(ovoNSQ.EventGive),
			},
		},
	})

	err = mq.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":10000", nil))
}
