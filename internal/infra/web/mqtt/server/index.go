package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
)

func (b *Broker) StartServer() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	opts := mqtt.Options{
		InlineClient: true,
	}

	server := mqtt.New(&opts)

	_ = server.AddHook(new(auth.Hook), &auth.Options{
		Ledger: &auth.Ledger{
			Auth: auth.AuthRules{
				{Username: auth.RString(b.username), Password: auth.RString(b.password), Allow: true},
			}}})

	listener := listeners.Config{
		Address: fmt.Sprintf("%s:%d", b.host, b.port),
		ID:      "t1",
	}

	tcp := listeners.NewTCP(listener)
	err := server.AddListener(tcp)

	if err != nil {
		panic(err)
	}

	callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
		fmt.Println("inline client received message from subscription", "client", cl.ID, "subscriptionId", sub.Identifier, "topic", pk.TopicName, "payload", string(pk.Payload))
	}
	server.Subscribe("topic/test", 1, callbackFn)

	go func() {
		err := server.Serve()
		if err != nil {
			panic(err)
		}
	}()

	sigReceived := <-sigs
	server.Log.Info("Received signal", "signal", sigReceived)
	server.Close()
}
