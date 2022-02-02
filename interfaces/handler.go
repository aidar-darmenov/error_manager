package interfaces

type Handler interface {
	Start()
	UseCase() UseCase
}
