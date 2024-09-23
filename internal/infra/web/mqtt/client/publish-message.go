package client

import "fmt"

func (b *Broker) PublishMessage(message string) {
	token := b.client.Publish(b.topic, 1, false, message)
	token.Wait()
	fmt.Printf("Published message to topic: %s\n", b.topic)
}
