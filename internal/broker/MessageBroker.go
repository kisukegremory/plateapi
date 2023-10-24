package broker

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		slog.Error("%s : %s", msg, err)
		panic("Shutdown the program due problems on consumer")
	}
}

func PublishJson(bindingKey string, body []byte) error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ChannelConnection.PublishWithContext(
		ctx,
		"",
		bindingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return fmt.Errorf("problems on publishing in the queue: %v", err)
	}
	return nil
}
