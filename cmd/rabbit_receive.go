package cmd

import (
	"github.com/orensimple/otus_hw1_8/config"
	"github.com/orensimple/otus_hw1_8/internal/logger"
	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

var RabbitRecieveCmd = &cobra.Command{
	Use:   "rabbit_recieve",
	Short: "Run rabbit recieve",
	Run: func(cmd *cobra.Command, args []string) {
		config.Init(addr)
		logger.InitLogger()
		startRecieve()
	},
}

func init() {
	RabbitRecieveCmd.Flags().StringVar(&addr, "config", "./config", "")
	RootCmd.AddCommand(RabbitRecieveCmd)
}

func startRecieve() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		logger.ContextLogger.Errorf("Failed to connect to RabbitMQ", err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.ContextLogger.Errorf("Failed to open a channel", err.Error())
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"eventsByDate", // name
		"direct",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		logger.ContextLogger.Infof("Failed to declare an exchange", err.Error())
	}

	_, err = ch.QueueDeclare(
		"eventsByDay", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	err = ch.QueueBind(
		"eventsByDay",  // name
		"day",          // key
		"eventsByDate", // exchange
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		logger.ContextLogger.Infof("Problem bind queue", err.Error())
	}

	msgs, err := ch.Consume(
		"eventsByDay",   // queue
		"ConsumerByDay", // consumer
		true,            // auto-ack
		false,           // exclusive
		false,           // no-local
		false,           // no-wait
		nil,             // args
	)
	if err != nil {
		logger.ContextLogger.Errorf("Failed to register a consumer", err.Error())
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			logger.ContextLogger.Infof("Received a message:", d.Body)
		}
	}()
	logger.ContextLogger.Infof(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
