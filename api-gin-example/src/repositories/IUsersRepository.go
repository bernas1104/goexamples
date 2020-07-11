package repositories

import "github.com/bernas1104/goexamples/api-gin-example/src/entities"

type UsersRepository interface {
	Save(user entities.User) entities.User
	Delete(id uint64)
	FindAll() []entities.User
	FindByID(id uint64) entities.User
	CloseDB()
}
