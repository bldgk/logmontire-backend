package rmq

import (
	"github.com/streadway/amqp"
	log2 "log"
	"logmontire/common/errors"
	"logmontire/common/functional"
	"logmontire/log"
)

const (
	CONNECTION_ERROR_CODE       = 700
	CHANNEL_ERROR_CODE          = 701
	DECLARE_EXCHANGE_ERROR_CODE = 702
	DECLARE_QUEUE_ERROR_CODE    = 703
	BIND_QUEUE_ERROR_CODE       = 704
)

type RabbitClient struct {
	connectionString string
	connection       *amqp.Connection
	channel          *amqp.Channel
}

func (rc *RabbitClient) init(connectionString string) {
	rc.connectionString = connectionString
}

func (rc *RabbitClient) connect() *errors.Error {
	connection, e := amqp.Dial(rc.connectionString)
	if e != nil {
		err := ConnectionError()
		log.Log(log.LStr(err.Error()), log.ConsoleLog)
		return err
	}

	rc.connection = connection
	return nil
}

func (rc *RabbitClient) createChannel() *errors.Error {
	channel, e := rc.connection.Channel()
	if e != nil {
		err := ChannelError()
		//logging.Log(logging.LStr(err.Code), logging.ConsoleLog)
		return err
	}

	rc.channel = channel
	return nil
}

func (rc *RabbitClient) declareExchange(exName string, exType string) *errors.Error {
	e := rc.channel.ExchangeDeclare(
		exName, // name
		exType, // type
		true,   // durable
		false,  // auto-deleted
		false,  // internal
		false,  // no-wait
		nil,    // arguments
	)
	if e != nil {
		err := DeclareExchangeError(exName, exType)
		log.Log(log.LStr(err.Error()), log.ConsoleLog)
		return err
	}

	return nil
}

func (rc *RabbitClient) makeQueue(queueName string) *errors.Error {
	_, e := rc.channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when usused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if e != nil {
		err := DeclareQueueError(queueName)
		log.Log(log.LStr(err.Error()), log.ConsoleLog)
		return err
	}

	return nil
}

func (rc *RabbitClient) bindQueue(exchange, queueName string, routingKey string) *errors.Error {
	e := rc.channel.QueueBind(
		queueName,  // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil,
	)
	if e != nil {
		err := BindQueueError(queueName, routingKey, exchange)
		log.Log(log.LStr(err.Error()), log.ConsoleLog)
		return err
	}

	return nil
}

func (rc *RabbitClient) createConsumer(queueName string, handler func([]byte)) {
	msgs, err := rc.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local3
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Log(log.LStr("Failed to register a consumer"), log.ConsoleLog)
	}

	go func() {
		for d := range msgs {
			handler(d.Body)
		}
	}()
}

func (rc *RabbitClient) close() {
	defer rc.channel.Close()
	defer rc.connection.Close()
}

var (
	rabbitClient = &RabbitClient{}
	//exchange     = "messenger"
)

func Init(url string) *errors.Error {
	rabbitClient.init(url)
	err := functional.Fs(rabbitClient.connect, rabbitClient.createChannel)
	if err != nil {
		log2.Fatal(err)
	}
	return nil
}

func Close() {
	rabbitClient.close()
}

func AddDirectExchange(name string) {
	rabbitClient.declareExchange(name, "direct")
}

func AddFanoutExchange(name string) {
	rabbitClient.declareExchange(name, "fanout")
}

//func publishMessage(routingKey string, body []byte, contentType string) {
//	err := channel.Publish(
//		exchange,   // exchange
//		routingKey, // routing key
//		false,      // mandatory
//		false,      // immediate
//		amqp.Publishing{
//			ContentType: contentType,
//			Body:        body,
//		})
//	if err != nil {
//		logging.Log(logging.LStr("Failed to publish a message"), logging.ConsoleLog)
//	}
//}
//
////
////type Message struct {
////	body        []byte
////	contentType string
////}
////
////func (m *Message) New(body []byte, contentType string) {
////	m.contentType = contentType
////	m.body = body
////}
////
////func (m *Message) Publish(routingKey string) {
////	PublishMessage(routingKey, m)
////}
//
////func NewMessage() *Message {
////	if contentType == "" {
////		contentType = "text/plain"
////	}
////	message := &Message{body, contentType}
////	return message
////}
//
//func PublishMessage(routingKey string, body []byte, contentType string) {
//	publishMessage(routingKey, body, contentType)
//}
//
func MakeQueue(queueName string) *errors.Error {
	return rabbitClient.makeQueue(queueName)
}

func BindQueue(exchange string, queueName string, routingKey string) *errors.Error {
	return rabbitClient.bindQueue(exchange, queueName, routingKey)
}

func CreateConsumer(queueName string, handler func([]byte)) {
	rabbitClient.createConsumer(queueName, handler)
}

//func MakeConsumer(queueName string) {
//
//	msgs, err := channel.Consume(
//		queueName, // queue
//		"",        // consumer
//		true,      // auto-ack
//		false,     // exclusive
//		false,     // no-local3
//		false,     // no-wait
//		nil,       // args
//	)
//	if err != nil {
//		logging.Log(logging.LStr("Failed to register a consumer"), logging.ConsoleLog)
//	}
//
//	go func() {
//		for d := range msgs {
//			fmt.Printf(" saved %s", d.Body)
//		}
//	}()
//}
