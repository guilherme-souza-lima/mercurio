package request

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Cellphone string `json:"cellphone"`
}

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Verify struct {
	ID        string `json:"id"`
	IDPoints  string `json:"uuid"`
	Email     string `json:"email"`
	Cellphone string `json:"cellphone"`
	Token     string `json:"token"`
}
