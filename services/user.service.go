package services

import "github.com/NeoTRAN001/go-ginframework/models"

/*
	Creamos nuestra interface para el servicio, este es un patron muy usado en MVC
	donde tenemos nuestra interface que luego será implementado por alguna struct que necesite
	interactuar con el repository.

	Aquí definimos los métodos que requerimos para funcionar
*/

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
