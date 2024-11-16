package internal

import (
	"errors"
	"time"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/TOMMy-Net/air-sell/models"
	"github.com/go-playground/validator/v10"
	datepicker "github.com/sdassow/fyne-datepicker"
)

var (
	ErrValidRow = errors.New("Не все поля были заполнены")
)

// Поле ввода даты
func (s *Settings) DateButton() *widget.Button {
	dateInput := widget.NewButton("", func() {})

	dateInput.OnTapped = func() {
		var d *dialog.CustomDialog

		when, err := time.Parse(dateInput.Text, "2006/01/02")
		if err != nil {
			when = time.Now()
		}

		datepicker := datepicker.NewDatePicker(when, time.Monday, func(when time.Time, ok bool) {
			if ok {
				dateInput.SetText(when.Format("2006/01/02"))
			}
			d.Hide()
		})

		d = dialog.NewCustomWithoutButtons("Выберите дату", datepicker, s.Window)
		d.Show()
	}
	return dateInput

}

// Поиск билетов
func (s *Settings) FindTickets(ticket *models.TicketsSearch) ([]models.Ticket, error) {
	validate := validator.New()
	err := validate.Struct(ticket)
	if err != nil {
		return []models.Ticket{}, ErrValidRow
	}
	tickets, err := s.Storage.FindTickets(ticket)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

// кнопка профиля
func (s *Settings) ProfileButton() *widget.Button{
	button := widget.NewButtonWithIcon("Профиль", theme.AccountIcon(), func() {
		
	})
	return button
}
