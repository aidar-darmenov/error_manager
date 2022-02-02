package usecase

func (u *UseCase) GetErrorById() {
	u.Repository.Mock().GetErrorById()
}
func (u *UseCase) AddError() {
	u.Repository.Mock().AddError()
}
func (u *UseCase) GetAllErrors() {
	u.Repository.Mock().GetAllErrors()
}
func (u *UseCase) DeleteError() {
	u.Repository.Mock().DeleteError()
}
func (u *UseCase) EditError() {
	u.Repository.Mock().EditError()
}
