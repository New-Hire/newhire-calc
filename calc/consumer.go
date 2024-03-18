package calc

import (
	"context"
	"fmt"
	rmqClient "github.com/apache/rocketmq-clients/golang/v5"
	log "github.com/sirupsen/logrus"
	"os"
	"time"

	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

const (
	Topic         = "newhire"
	ConsumerGroup = "HIRENEW"
	Endpoint      = "localhost:8080"
	AccessKey     = ""
	SecretKey     = ""
)

var (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// maximum number of messages received at one time
	maxMessageNum int32 = 1
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// receive messages in a loop
)

func Consumer() {
	// log to console
	os.Setenv("mq.consoleAppender.enabled", "true")
	rmqClient.ResetLogger()
	// In most case, you don't need to create many consumers, singleton pattern is more recommended.
	simpleConsumer, err := rmqClient.NewSimpleConsumer(&rmqClient.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmqClient.WithAwaitDuration(awaitDuration),
		rmqClient.WithSubscriptionExpressions(map[string]*rmqClient.FilterExpression{
			Topic: rmqClient.SUB_ALL,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = simpleConsumer.Start()
	if err != nil {
		log.Fatal(err)
	}
	defer simpleConsumer.GracefulStop()

	go func() {
		for {
			fmt.Println("start receive message")
			mvs, err := simpleConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
			if err != nil {
				fmt.Println(err)
			}
			for _, mv := range mvs {
				err := simpleConsumer.Ack(context.TODO(), mv)
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(mv)
				fmt.Println(string(mv.GetBody()))
			}
			fmt.Println("wait a moment")
			fmt.Println()
			time.Sleep(time.Second * 3)
		}
	}()
	// run for a while
	time.Sleep(time.Minute)
}
