package request

import "injar/usecase/users"

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		ID:       req.ID,
		Name:     req.Name,
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	}
}
