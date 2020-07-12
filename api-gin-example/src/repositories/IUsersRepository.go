package repositories

import "github.com/bernas1104/goexamples/api-gin-example/src/entities"

type UsersRepository interface {
	Create(user entities.User) (entities.User, error)
	Delete(id uint64)
	FindAll() []entities.User
	FindByEmail(email string) entities.User
	FindByID(id uint64) entities.User
	CloseDB()
}
