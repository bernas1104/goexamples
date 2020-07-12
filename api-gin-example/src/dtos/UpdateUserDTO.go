package dtos

type UpdateUserDTO struct {
	Name                    string `json:"name" binding:"required,max=100"`
	Age                     int8   `json:"age" binding:"gte=1,lte=100"`
	Email                   string `json:"email" binding:"required,email"`
	OldPassword             string `json:"old_password" binding:"omitempty"`
	NewPassword             string `json:"new_password" binding:"omitempty,required_with=OldPassword,min=6,max=12" json:"-"`
	NewPasswordConfirmation string `json:"new_password_confirmation" binding:"required_with=OldPassword,eqfield=NewPassword" json:"-"`
}
