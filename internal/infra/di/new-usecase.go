package di

import (
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/cassandra/repository"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/factory"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
)

func NewUseCase(db *factory.Iconnection) *usecase.UseCaseMessageUser {

	if db.Gocql != nil {
		repository := repository.NewCassandraRepository(db)
		return usecase.NewUsecase(repository)
	}

	return nil
}
