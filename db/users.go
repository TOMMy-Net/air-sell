package db

import (
	"github.com/TOMMy-Net/air-sell/models"
	"github.com/TOMMy-Net/air-sell/tools"
)

func (s *Storage) CreateUser(u models.User) (int64, error) {
	u.Password = tools.Sum256([]byte(u.Password))
	res, err := s.DB.NamedExec(`INSERT INTO users(email, password) VALUES (:email, :password)`, u)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}

func (s *Storage) GetUser(u models.User) (*models.User, error) {
	user := models.User{}
	u.Password = tools.Sum256([]byte(u.Password))
	err := s.DB.Get(&user, `SELECT * FROM users WHERE email = $1 AND password = $2`, u.Email, u.Password)
	return &user, err
}

func (s *Storage) SetBuyHistory(h models.BuyHistory) error {
	_, err := s.DB.Exec(`INSERT INTO buy_history(ticket_id, user_id, buy_time, count) VALUES ($1, $2, $3, $4)`, h.Ticket.ID, h.UserId, h.BuyTime, h.Count)
	return err
}
