package response

type UserLogin struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
	Token     string `json:"token"`
	Points    Points `json:"points"`
}

type Points struct {
	ID       string `json:"uuid"`
	GGPoints int    `json:"gg_points"`
}
