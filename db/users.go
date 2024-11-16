package db

import (

	"github.com/TOMMy-Net/air-sell/models"
	"github.com/TOMMy-Net/air-sell/tools"
)

func (s *Storage) CreateUser(u models.User) error {
	u.Password = tools.Sum256([]byte(u.Password))
	_, err := s.DB.NamedExec(`INSERT INTO users(email, password) VALUES (:email, :password)`, u)
	return err
}

func (s *Storage) GetUser(u models.User) (*models.User, error) {
	user := models.User{}
	u.Password = tools.Sum256([]byte(u.Password))
	err := s.DB.Get(&user, `SELECT * FROM users WHERE email = $1 AND password = $2`, u.Email, u.Password)
	return &user, err
}
