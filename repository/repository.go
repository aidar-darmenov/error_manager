package repository

import (
	"github.com/aidar-darmenov/error_manager/interfaces"
	"github.com/aidar-darmenov/error_manager/repository/mock"
)

type Repository struct {
	RMock interfaces.Mock
	//DB
}

func NewRepository() interfaces.Repository {
	return &Repository{RMock: &mock.Mock{}}
}

func (r *Repository) Mock() interfaces.Mock {
	return r.RMock
}
