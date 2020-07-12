package dtos

type CreateUserDTO struct {
	Name                 string `json:"name" binding:"required,max=100"`
	Age                  int8   `json:"age" binding:"gte=1,lte=100"`
	Email                string `json:"email" binding:"required,email"`
	Password             string `json:"password" binding:"required,min=6,max=12"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required,eqfield=Password" json:"-"`
}
