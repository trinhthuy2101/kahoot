package user

import (
	"examples/identity/jwthelper"
	"examples/identity/model/user"
	model "examples/identity/model/user"
)

type UserHandler interface {
	Login(request *BaseUserRequest) (string, error)
	Register(request *BaseUserRequest) (string, error)
}

type handler struct {
	repo       model.Repository
	jwtService jwthelper.JWTHelper
}

func (user *handler) Login(request *BaseUserRequest) (string, error) {
	isUserAuthenticated := user.repo.Login(&model.User{ID: request.ID, Username: request.Username, Password: request.Password})

	var token string
	var err error = nil

	if isUserAuthenticated {
		token, err = user.jwtService.GenerateJWT(request.Username)
		return token, err
	}
	return "", err
}
func (user *handler) Register(request *BaseUserRequest) (string, error) {
	isUserAuthenticated := user.repo.Register(&model.User{ID: request.ID, Username: request.Username, Password: request.Password})

	var token string
	var err error = nil

	if isUserAuthenticated {
		token, err = user.jwtService.GenerateJWT(request.Username)
		return token, err
	}
	return "", err
}
func NewHandler(repo user.Repository, jwt_service jwthelper.JWTHelper) UserHandler {
	return &handler{
		repo:       repo,
		jwtService: jwt_service,
	}
}
