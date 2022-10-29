package services

import (
	"context"

	"github.com/NeoTRAN001/go-ginframework/models"
	"github.com/NeoTRAN001/go-ginframework/repositories"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
	Implementamos la interface dentro de nuestra estruct, en el contructor NewUserService,
	y dentro hacemos una inyección de dependencias del repository, que será el encargado de hacer
	las query a la DB
*/

type UserServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewUserService(
	mongoClient *mongo.Client,
	ctxDB context.Context,
) UserService {
	return &UserServiceImpl{
		userRepository: repositories.NewUserRepository(mongoClient, ctxDB),
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	return u.userRepository.Create(user)
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	return u.userRepository.GetUser(name)
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	return u.userRepository.GetAll()
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	return u.userRepository.UpdateUser(user)
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	return u.userRepository.DeleteUser(name)
}
