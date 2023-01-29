package client

import (
	rmq "github.com/rabbitmq/amqp091-go"
	"context"
	"fmt"
	"time"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func connectRMQ() (*rmq.Connection, *rmq.Channel)  {

	connection, err := rmq.Dial("amqp://guest:guest@localhost:5672/")
	checkErr(err)
	defer connection.Close() //clsoe conenciton at end of func 	

	channel, err := connection.Channel()
	checkErr(err)
	defer channel.Close() //close channel at end of func
	//channel is an access point to the rabbiutmqw connectio

	//idempotent queue declare (doenst exist wilklc reate)

	return connection, channel

}

func sendMessage(con *rmq.Connection, cha *rmq.Channel, message string, senderId string, receiverId string) {
	queue, err := cha.QueueDeclare(senderId + "_" + receiverId, false, false, false, false, nil)
	checkErr(err)
	//idempotent queue declare (doenst exist wilklc reate)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = cha.PublishWithContext(ctx, "", queue.Name, false, false, rmq.Publishing{ContentType: "text/plain", Body: []byte(message)})

	checkErr(err)
	fmt.Printf(" [x] Sent %s\n", message)
}

