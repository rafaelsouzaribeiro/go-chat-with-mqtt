package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
)

func (i *CassandraRepository) PublishMessage(input *entity.Message) error {

	insert := fmt.Sprintf("INSERT INTO %s.messages(userid,message,username,pages,times)VALUES(?,?,?,?,?)", entity.KeySpace)
	err := i.gocql.Query(insert, uuid.NewString(), input.Message, input.Username, 1, time.Now()).Exec()

	if err != nil {
		return err
	}

	return nil
}
