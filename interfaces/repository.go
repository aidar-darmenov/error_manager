package interfaces

type Repository interface {
	Mock() Mock
}

type Mock interface {
	GetErrorById()
	AddError()
	GetAllErrors()
	DeleteError()
	EditError()
}
