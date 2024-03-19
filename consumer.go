package main

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
	Topic         = "test"
	ConsumerGroup = "TEST"
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

func consumer(method func(body string)) {
	err := os.Setenv("mq.consoleAppender.enabled", "true")
	if err != nil {
		panic("[MQ] 设置环境变量异常")
	}
	rmqClient.ResetLogger()
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

	defer func(simpleConsumer rmqClient.SimpleConsumer) {
		err := simpleConsumer.GracefulStop()
		if err != nil {
			log.Fatal(err)
		}
	}(simpleConsumer)

	for {
		fmt.Println("开始新一轮消息接收")
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
			method(string(mv.GetBody()))
		}
		// TODO: 配置化秒数
		fmt.Println("等待2秒继续接收")
		time.Sleep(time.Second * 2)
	}
}
