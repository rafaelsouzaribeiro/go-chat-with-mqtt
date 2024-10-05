package repository

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/go-chat-with-mqtt/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

func (r *CassandraRepository) Login(username, password string) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	s := fmt.Sprintf(`SELECT id,username,photo FROM %s.users 
					 WHERE username=? AND password=?;`, entity.KeySpace)

	query := r.gocql.Query(s, username, hashedPassword)
	iter := query.Iter()
	defer iter.Close()

	var user entity.User

	for iter.Scan(&user.Id, &user.Username,
		&user.Photo) {

		return &user, nil
	}

	return nil, fmt.Errorf("user not found or incorrect credentials")
}
