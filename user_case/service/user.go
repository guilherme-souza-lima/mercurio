package service

import (
	"ssMercurio/entities"
	"ssMercurio/infrastructure/jwt"
	"ssMercurio/user_case/repository"
	"ssMercurio/user_case/request"
	"ssMercurio/user_case/response"
)

type CryptoPassword interface {
	Encrypt(password string) (string, error)
	Decrypt(crypt string) (string, error)
}

type UserService struct {
	UserRepository repository.RepositoryUser
	JwtToken       jwt.JwtToken
	CryptoPassword CryptoPassword
}

func NewUserService(UserRepository repository.RepositoryUser, JwtToken jwt.JwtToken, CryptoPassword CryptoPassword) UserService {
	return UserService{UserRepository, JwtToken, CryptoPassword}
}

func (u UserService) Create(data request.User) error {
	var entity entities.User
	result := entity.ToDomain(data)

	newPassword, err := u.CryptoPassword.Encrypt(data.Password)
	if err != nil {
		return err
	}
	result.Password = newPassword
	return u.UserRepository.Create(result)
}

func (u UserService) Login(data request.Login) (response.UserLogin, error) {
	var entity entities.Login
	var login response.UserLogin
	result := entity.ToDomain(data)

	newPassword, errCrypto := u.CryptoPassword.Encrypt(data.Password)
	if errCrypto != nil {
		return login, errCrypto
	}
	result.Password = newPassword

	user, err := u.UserRepository.Login(result)
	if err != nil {
		return login, err
	}
	login.ID = user.ID
	login.Email = user.Email
	login.Cellphone = user.Cellphone
	login.Points.ID = user.Points.ID
	login.Points.GGPoints = user.Points.GGPoints
	token, errToken := u.JwtToken.Create(login.ID, login.Email, login.Cellphone, login.Points.ID)
	if err != nil {
		return login, errToken
	}
	login.Token = token
	return login, nil
}

func (u UserService) Verify(data request.Verify) (bool, error) {
	result, err := u.JwtToken.Validation(data.Token)
	if err != nil {
		return false, err
	}
	if data.ID == result.ID &&
		data.Email == result.Email &&
		data.Cellphone == result.Cellphone &&
		data.IDPoints == result.IDPoints {
		return true, nil
	}
	return false, nil
}
