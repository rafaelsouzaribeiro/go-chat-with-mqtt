package server

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
	"github.com/mochi-mqtt/server/v2/packets"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase/dto"
	"github.com/spf13/viper"
)

func (b *Broker) StartServer() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	viper.AutomaticEnv()

	opts := mqtt.Options{
		InlineClient: true,
	}

	server := mqtt.New(&opts)

	_ = server.AddHook(new(auth.Hook), &auth.Options{
		Ledger: &auth.Ledger{
			Auth: auth.AuthRules{
				{Username: auth.RString(b.Username), Password: auth.RString(b.Password), Allow: true},
			}}})

	tcp := listeners.NewTCP(GetListeners(b.Host, "tcp", b.Port))
	err := server.AddListener(tcp)

	if err != nil {
		panic(err)
	}

	ws := listeners.NewWebsocket(GetListeners(b.SocketHost, "socket", b.SocketPort))
	err = server.AddListener(ws)
	if err != nil {
		panic(err)
	}

	callbackFn := func(cl *mqtt.Client, sub packets.Subscription, pk packets.Packet) {
		var Payload dto.Payload

		err := json.Unmarshal(pk.Payload, &Payload)
		if err != nil {
			fmt.Printf("Failed to deserialize message: %v\n", err)
			return
		}

		dto := dto.Payload{
			Message:  Payload.Message,
			Username: Payload.Username,
			UserId:   Payload.UserId,
		}

		_, err = b.Usecase.SaveMessage(&dto)

		if err != nil {
			fmt.Printf("Failed to save message: %v\n", err)
			return
		}
	}

	server.Subscribe(viper.GetString("TOPIC_MQTT"), 1, callbackFn)

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
