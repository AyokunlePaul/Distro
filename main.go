package main

import "github.com/streadway/amqp"

func main() {
	go client()
	server()
}

func server() {
	amqpConnection, amqpChannel, amqpQueue := getQueue()
	defer amqpConnection.Close()
	defer amqpChannel.Close()

	message := amqp.Publishing{
		Body:        []byte("Hello, Distro"),
		ContentType: "text/plain",
	}
	for {
		publishError := amqpChannel.Publish("", amqpQueue.Name, false, false, message)
		if publishError != nil {
			panic(publishError)
		}
	}
}

func client() {
	amqpConnection, amqpChannel, amqpQueue := getQueue()
	defer amqpConnection.Close()
	defer amqpChannel.Close()

	messages, messageError := amqpChannel.Consume(amqpQueue.Name, "", true, false,
		false, false, nil)
	if messageError != nil {
		panic(messageError)
	}
	for message := range messages {
		println(string(message.Body))
	}
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	amqpConnection, connectionError := amqp.Dial("amqp://guest@localhost:5672")
	if connectionError != nil {
		panic(connectionError)
	}
	amqpChannel, channelError := amqpConnection.Channel()
	if channelError != nil {
		panic(channelError)
	}
	amqpQueue, queueError := amqpChannel.QueueDeclare(
		"hello", false, false, false, false, nil,
	)
	if queueError != nil {
		panic(queueError)
	}
	return amqpConnection, amqpChannel, &amqpQueue
}
