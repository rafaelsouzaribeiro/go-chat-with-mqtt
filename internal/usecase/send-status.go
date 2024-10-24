package usecase

func (i *UseCaseMessageUser) SendStatus(id, status string) {

	i.Irepository.SendStatus(id, status)

}
