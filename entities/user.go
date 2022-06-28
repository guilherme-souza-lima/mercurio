package entities

import (
	"github.com/gofrs/uuid"
	"ssMercurio/user_case/request"
)

type User struct {
	ID        string `json:"id"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
	Points    Points `json:"points"`
}

type Points struct {
	ID       string `json:"uuid"`
	GGPoints int    `json:"gg_points"`
}

func (u User) ToDomain(data request.User) User {
	uuidGenerator, _ := uuid.NewV4()
	uuidPoints, _ := uuid.NewV4()
	return User{
		ID:        uuidGenerator.String(),
		Email:     data.Email,
		Cellphone: data.Cellphone,
		Points: Points{
			ID:       uuidPoints.String(),
			GGPoints: 10,
		},
	}
}

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (l Login) ToDomain(data request.Login) Login {
	return Login{
		Login: data.Login,
	}
}
