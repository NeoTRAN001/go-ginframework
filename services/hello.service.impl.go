package services

import "github.com/NeoTRAN001/go-ginframework/models"

type HelloServiceImpl struct {
}

func NewHelloService() HelloService {
	return &HelloServiceImpl{}
}

func (s *HelloServiceImpl) GetHello() (*models.Message, error) {
	message := models.Message{
		Title: "Hello",
		Body:  "World",
	}

	return &message, nil
}
