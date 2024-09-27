package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func (b *Broker) StartServer() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	opts := mqtt.Options{
		InlineClient: true,
	}

	b.Server = mqtt.New(&opts)

	_ = b.Server.AddHook(new(auth.Hook), &auth.Options{
		Ledger: &auth.Ledger{
			Auth: auth.AuthRules{
				{Username: auth.RString(b.username), Password: auth.RString(b.password), Allow: true},
			}}})

	listener := listeners.Config{
		Address: fmt.Sprintf("%s:%d", b.host, b.port),
		ID:      "t1",
	}

	tcp := listeners.NewTCP(listener)
	err := b.Server.AddListener(tcp)

	if err != nil {
		panic(err)
	}

	go func() {
		err := b.Server.Serve()
		if err != nil {
			panic(err)
		}
	}()

	sigReceived := <-sigs
	b.Server.Log.Info("Received signal", "signal", sigReceived)
	b.Server.Close()
}
