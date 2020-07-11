package entities

import "time"

// User - Defines the entity for an user with the fields ID, Name, Age and Email
type User struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Name      string    `json:"name" binding:"required,max=100" gorm:"type:varchar(100)"`
	Age       int8      `json:"age" binding:"gte=1,lte=100"`
	Email     string    `json:"email" binding:"required,email" gorm:"type:varchar(256)"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" form:"default:CURRENT_TIMESTAMP"`
}
