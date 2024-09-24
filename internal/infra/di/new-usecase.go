package di

import (
	"github.com/gocql/gocql"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/infra/database/cassandra/repository"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/usecase"
)

func NewUseCase(db *gocql.Session) *usecase.UseCaseMessageUser {

	repository := repository.NewCassandraRepository(db)
	return usecase.NewUsecase(repository)
}
