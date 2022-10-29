package repositories

import (
	"context"
	"errors"

	"github.com/NeoTRAN001/go-ginframework/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
 Las interface son la estructura que debe de tener los struct, que va a emular las class
 que tenemos por ejemplo en Java, así que un struc va a implementar la interface.
 Con la función NewUserRepository vamos a simular un constructor para nuestro "objeto"
*/

type UserRepository interface {
	Create(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}

type UserRepositoryImpl struct {
	collection  *mongo.Collection
	ctx         context.Context
	mongoClient *mongo.Client
}

func NewUserRepository(mongoClient *mongo.Client, ctx context.Context) UserRepository {
	return &UserRepositoryImpl{
		collection: mongoClient.Database("userdb").Collection("users"),
		ctx:        ctx,
	}
}

func (u *UserRepositoryImpl) Create(user *models.User) error {
	_, err := u.collection.InsertOne(u.ctx, user)
	return err
}

func (u *UserRepositoryImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User

	err := u.collection.FindOne(
		u.ctx,
		bson.D{bson.E{Key: "user_name", Value: name}},
	).Decode(&user)

	return user, err
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	var users []*models.User

	cursor, err := u.collection.Find(u.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(u.ctx) {
		var user models.User

		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(u.ctx)

	if len(users) == 0 {
		return nil, errors.New("Documents not found")
	}

	return users, nil
}

func (u *UserRepositoryImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{
		bson.E{
			Key: "$set", Value: bson.D{
				bson.E{Key: "user_name", Value: user.Name},
				bson.E{Key: "user_addres", Value: user.Address},
				bson.E{Key: "user_age", Value: user.Age},
			},
		},
	}

	result, _ := u.collection.UpdateOne(u.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("No matched document found for update")
	}

	return nil
}

func (u *UserRepositoryImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, _ := u.collection.DeleteOne(u.ctx, filter)

	if result.DeletedCount != 1 {
		return errors.New("No matched document found for delete")
	}

	return nil
}
