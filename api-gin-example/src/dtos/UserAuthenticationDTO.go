package dtos

type UserAuthenticationDTO struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"min=6,max=12,required"`
}
