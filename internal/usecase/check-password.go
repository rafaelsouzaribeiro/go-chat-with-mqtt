package usecase

import "golang.org/x/crypto/bcrypt"

func (i *UseCaseMessageUser) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
