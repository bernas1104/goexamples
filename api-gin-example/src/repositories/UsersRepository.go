package repositories

import (
	"fmt"

	"github.com/bernas1104/goexamples/api-gin-example/src/entities"
	"github.com/jinzhu/gorm"

	// GORM Dialect
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

func (repository *usersRepository) Create(user entities.User) (entities.User, error) {
	email := user.Email

	var userExists entities.User
	repository.ormRepository.Where(&entities.User{Email: email}).First(&userExists)

	if userExists.ID != 0 {
		return entities.User{}, fmt.Errorf("E-mail is already being used")
	}

	repository.ormRepository.Create(&user)

	return user, nil
}

func (repository *usersRepository) Save(user entities.User) entities.User {
	repository.ormRepository.Save(&user)

	return user
}

func (repository *usersRepository) Delete(email string) {
	var deleteUser entities.User

	repository.ormRepository.Where(&entities.User{Email: email}).Delete(&deleteUser)
}

func (repository *usersRepository) FindAll() []entities.User {
	var users []entities.User

	repository.ormRepository.Find(&users)

	return users
}

func (repository *usersRepository) FindByEmail(email string) entities.User {
	var user entities.User

	repository.ormRepository.Where(&entities.User{Email: email}).First(&user)

	return user
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
