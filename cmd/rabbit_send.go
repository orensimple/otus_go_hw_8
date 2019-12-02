package cmd

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/orensimple/otus_hw1_8/config"
	"github.com/orensimple/otus_hw1_8/internal/domain/services"
	"github.com/orensimple/otus_hw1_8/internal/logger"
	"github.com/orensimple/otus_hw1_8/internal/maindb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

var RabbitSendCmd = &cobra.Command{
	Use:   "rabbit_send",
	Short: "Run rabbit send",
	Run: func(cmd *cobra.Command, args []string) {
		config.Init(addr)
		logger.InitLogger()
		startSend()
	},
}

func init() {
	RabbitSendCmd.Flags().StringVar(&addr, "config", "./config", "")
	RootCmd.AddCommand(RabbitSendCmd)
}

func startSend() {
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
	if err != nil {
		logger.ContextLogger.Infof("Problem connect to", dsn, err.Error())
	}

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

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/events", viper.GetString("postgres.user"), viper.GetString("postgres.passwd"), viper.GetString("postgres.ip"), viper.GetString("postgres.port"))

	eventStorage, err := maindb.NewPgEventStorage(dsn)

	if err != nil {
		logger.ContextLogger.Errorf("Problem connect to db", dsn, err.Error())
	}

	eventService := &services.EventService{
		EventStorage: eventStorage,
	}
	ctx := context.Background()
	logger.ContextLogger.Infof(" [*] Start to check events. To exit press CTRL+C")
	uptimeTicker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-uptimeTicker.C:
			sendEvents(ctx, ch, eventService)
		}
	}

}

func sendEvents(ctx context.Context, ch *amqp.Channel, eventService *services.EventService) {
	events, _ := eventService.GetEventsByTime(ctx, "day")
	for _, event := range events {
		err := ch.Publish(
			"eventsByDate", // exchange
			"day",          // routing key
			false,          // mandatory
			false,          // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(strconv.FormatInt(event.ID, 10)),
			})

		if err != nil {
			logger.ContextLogger.Errorf("Failed to publish a message", err.Error())
		}
		logger.ContextLogger.Infof(" Publish message", event.ID)
	}
}
