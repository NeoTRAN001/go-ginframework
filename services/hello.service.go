package services

import "github.com/NeoTRAN001/go-ginframework/models"

type HelloService interface {
	GetHello() (*models.Message, error)
}
