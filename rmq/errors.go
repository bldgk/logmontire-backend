package rmq

import (
	"fmt"
	"logmontire/common/errors"
)

func ConnectionError() *errors.Error {
	return &errors.Error{
		Code: CONNECTION_ERROR_CODE,
		What: fmt.Sprintf("Failed to connect to RabbitMQ"),
	}
}

func ChannelError() *errors.Error {
	return &errors.Error{
		Code: CHANNEL_ERROR_CODE,
		What: fmt.Sprintf("Failed to open a channel"),
	}
}

func DeclareExchangeError(exName string, exType string) *errors.Error {
	return &errors.Error{
		Code: DECLARE_EXCHANGE_ERROR_CODE,
		What: fmt.Sprintf("Failed to declare an exchange %s of type %s", exName, exType),
	}
}

func DeclareQueueError(queueName string) *errors.Error {
	return &errors.Error{
		Code: DECLARE_QUEUE_ERROR_CODE,
		What: fmt.Sprintf("Failed to declare a queue %s", queueName),
	}
}

func BindQueueError(queueName string, routingKey string, exchange string) *errors.Error {
	return &errors.Error{
		Code: BIND_QUEUE_ERROR_CODE,
		What: fmt.Sprintf("Failed to bind a queue %s to %s in %s", queueName, routingKey, exchange),
	}
}
