package di

import (
	"strings"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/cassandra/connection"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/cassandra/repository"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
	"github.com/spf13/viper"
)

func NewUseCase() *usecase.UseCaseMessageUser {
	viper.AutomaticEnv()
	hosts := strings.Split(viper.GetString("HOST_CASSANDRA"), ",")
	db, errs := connection.NewCassandraConnect(hosts, viper.GetString("USER_CASSANDRA"), viper.GetString("PASSWORD_CASSANDRA"))

	if errs != nil {
		panic(errs)
	}
	repository := repository.NewCassandraRepository(db)
	return usecase.NewUsecase(repository)
}
