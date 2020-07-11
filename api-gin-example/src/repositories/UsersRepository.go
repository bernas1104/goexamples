package repositories

import (
	"github.com/bernas1104/goexamples/api-gin-example/src/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type usersRepository struct {
	ormRepository *gorm.DB
}

func NewUsersRepository() UsersRepository {
	db, err := gorm.Open("sqlite3", "test.db") // Connects to the database

	if err != nil {
		panic("Failed to connect to the database") // Application can't work without DB
	}

	db.AutoMigrate(&entities.User{})

	return &usersRepository{
		ormRepository: db,
	}
}

func (repository *usersRepository) Save(user entities.User) entities.User {
	id := user.ID

	var userExists entities.User
	repository.ormRepository.Where(&entities.User{ID: id}).First(&userExists)

	if userExists.ID != 0 {
		repository.ormRepository.Save(&user)
	} else {
		repository.ormRepository.Create(&user)
	}

	return user
}

func (repository *usersRepository) Delete(id uint64) {
	var deleteUser entities.User

	repository.ormRepository.Where(&entities.User{ID: id}).Delete(&deleteUser)
}

func (repository *usersRepository) FindAll() []entities.User {
	var users []entities.User

	repository.ormRepository.Find(&users)

	return users
}

func (repository *usersRepository) FindByID(id uint64) entities.User {
	var user entities.User

	repository.ormRepository.Where(&entities.User{ID: id}).First(&user)

	return user
}

func (repository *usersRepository) CloseDB() {
	err := repository.ormRepository.Close()

	if err != nil {
		panic("Failed to close connection to the database")
	}
}
