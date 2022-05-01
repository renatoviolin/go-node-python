package ports

import "github.com/renatoviolin/go-node/dto"

type Repository interface {
	FindAll() (*[]dto.Payload, error)
}
