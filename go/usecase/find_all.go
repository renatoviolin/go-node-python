package usecase

import (
	"github.com/renatoviolin/go-node/dto"
	"github.com/renatoviolin/go-node/ports"
)

type FindAll struct {
	repository ports.Repository
}

func NewFindAllUseCase(repository ports.Repository) *FindAll {
	return &FindAll{repository: repository}
}

func (f *FindAll) Execute() (*[]dto.Payload, error) {
	result, err := f.repository.FindAll()
	return result, err
}
