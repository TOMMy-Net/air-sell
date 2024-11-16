package models

type TicketsSearch struct {
	From      string `validate:"required,min=1"`
	To        string `validate:"required,min=1"`
	Date_from string `validate:"required,min=1"`
	Date_to   string `validate:"required,min=1"`
}

type AirPorts struct {
	ID      int    `db:"id"`
	City    string `db:"city"`
	Iata    string `db:"iata_code"`
	Country string `db:"country"`
}

type Ticket struct {
	ID            string   `db:"id"`             // UUID
	Airline       string   `db:"airline"`        // рейс
	DepartureFrom AirPorts `db:"departure_from"` // город
	ArrivalAt     AirPorts `db:"arrival_at"`     // город
	DepartureTime string   `db:"departure_time"` // дата
	ArrivalTime   string   `db:"arrival_time"`   // дата
	Quantity      int      `db:"quantity"`       // колличество
	Luggage       string   `db:"luggage"`        // багаж
	HandBaggage   string   `db:"hand_baggage"`   // ручная кладь
	Price         float64  `db:"price"`          // цена
}

type User struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type Passport struct {
	ID              int    `db:"id"`
	Name            string `db:"name"`
	Surname         string `db:"surname"`
	Patronymic      string `db:"patronymic"`
	SeriesAndNumber string `db:"passport_series_and_number"`
	Gender          string `db:"gender"`
	ValidityPeriod  string `db:"validity_period"`
	Birthday        string `db:"date_of_birth"`
	Type            string `db:"passport_type"`
	Citizenship     string `db:"citizenship"`
	UserID          int    `db:"user_id"`
}

type UserEntry struct {
	Email    string `validate:"required,min=1,email"`
	Password string `validate:"required,min=1"`
}

// Функция создания структуры билета
func NewTicketSearch() *TicketsSearch {
	return new(TicketsSearch)
}
