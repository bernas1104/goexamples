package services

type UserAuthenticationService interface {
	AuthenticateUser(username string, password string) bool
}

type userAuthenticationService struct {
	authorizedUsername string // Placeholder; No-DB yet
	authorizedPassword string // Placeholder; No-DB yet
}

func NewUserAuthenticationService() UserAuthenticationService {
	return &userAuthenticationService{
		authorizedUsername: "johndoe@example.com",
		authorizedPassword: "123456",
	}
}

func (service *userAuthenticationService) AuthenticateUser(
	username string,
	password string,
) bool {
	return service.authorizedUsername == username &&
		service.authorizedPassword == password
}
