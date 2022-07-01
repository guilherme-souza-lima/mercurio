package service

import (
	"errors"
	"ssMercurio/infrastructure/jwt"
	"ssMercurio/user_case/repository"
	"ssMercurio/user_case/request"
)

type PointsService struct {
	JwtToken       jwt.JwtToken
	UserRepository repository.RepositoryUser
}

func NewPointsService(JwtToken jwt.JwtToken, UserRepository repository.RepositoryUser) PointsService {
	return PointsService{JwtToken, UserRepository}
}

func (p PointsService) VerifyPoints(data request.PointsRequest) (bool, error) {
	result, err := p.JwtToken.Validation(data.Verify.Token)
	if err != nil {
		return false, err
	}
	if data.Verify.ID == result.ID &&
		data.Verify.Email == result.Email &&
		data.Verify.Cellphone == result.Cellphone &&
		data.Verify.IDPoints == result.IDPoints {
		user, _ := p.UserRepository.GetUserById(data.Verify.ID)
		point, errPoint := calculation(data.Points.Value, user.Points.GGPoints)
		if errPoint != nil {
			return false, errPoint
		}
		update := p.UserRepository.UpdatePointsUser(data.Verify.ID, point)
		if update != nil {
			return false, update
		}
		return true, nil
	}
	return false, nil
}

func calculation(n1, n2 int) (int, error) {
	if n1 <= n2 {
		return n2 - n1, nil
	}
	return 0, errors.New("insufficient points")
}
