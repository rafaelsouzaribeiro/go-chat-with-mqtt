package usecase

func (i *UseCaseMessageUser) GetUsername(username string) int {
	u := i.Irepository.GetUsername(username)

	return u
}
