package internal

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"github.com/TOMMy-Net/air-sell/db"
	"github.com/TOMMy-Net/air-sell/models"
	"github.com/TOMMy-Net/air-sell/tools"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2/widget"
)

type Settings struct {
	Window  fyne.Window
	Storage *db.Storage
	Account *models.User
}

// Функция создания структуры настроек
func NewSettings() *Settings {
	return &Settings{}
}

func (s *Settings) SignInWindow() {
	email := widget.NewEntry()
	password := widget.NewPasswordEntry()
	form := widget.NewForm(
		widget.NewFormItem("Почта", email),
		widget.NewFormItem("Пароль", password),
	)

	form.SubmitText = "Войти"
	form.Orientation = widget.Vertical

	form.OnSubmit = func() {
		m := models.UserEntry{
			Email:    email.Text,
			Password: password.Text,
		}
		if err := tools.Validate(m); err != nil {
			dialog.ShowError(err, s.Window)
		} else {
			user, err := s.Storage.GetUser(models.User{Email: m.Email, Password: m.Password})
			if err != nil {
				dialog.ShowInformation("Ошибка", "Ошибка базы данных", s.Window)
			} else {
				s.Account = user
				s.MainWindow()
			}
		}
	}

	label := widget.NewLabel("для того чтобы продолжить войдите в систему")

	fTxt := canvas.NewText("Добро Пожаловать", color.White)
	fTxt.TextSize = 20
	fTxt.TextStyle.Bold = true
	fTxt.Alignment = fyne.TextAlignCenter

	regButton := widget.NewButtonWithIcon("Зарегестрироваться", theme.AccountIcon(), func() {
		s.SignUpWindow()
	})

	withOutReg := widget.NewButton("Продолжить без регистрации", func() {
		s.MainWindow()
	})
	regButton.Enable()

	s.Window.SetContent(
		container.NewBorder(nil, nil, nil, nil,
			container.NewCenter(container.NewVBox(fTxt, label, canvas.NewLine(color.White), form, widget.NewActivity(), regButton, withOutReg))))

}

func (s *Settings) SignUpWindow() {
	email := widget.NewEntry()
	password := widget.NewPasswordEntry()
	confirmPassword := widget.NewPasswordEntry()
	form := widget.NewForm(
		widget.NewFormItem("Почта", email),
		widget.NewFormItem("Пароль", password),
		widget.NewFormItem("Повторите пароль", confirmPassword),
	)

	form.SubmitText = "Зарегестрироваться"
	form.Orientation = widget.Vertical

	form.OnSubmit = func() {
		if password.Text == confirmPassword.Text {
			m := models.UserEntry{
				Email:    email.Text,
				Password: password.Text,
			}
			if err := tools.Validate(m); err != nil {
				dialog.ShowInformation("Ошибка", "Не все поля заполнены", s.Window)
			} else {
				if err := s.Storage.CreateUser(models.User{Email: m.Email, Password: m.Password}); err != nil {
					dialog.ShowInformation("Ошибка", "Не правильные почта или пароль", s.Window)
				} else {
					s.MainWindow()
				}
			}
		} else {
			dialog.ShowInformation("Ошибка", "Пароли не совпадают", s.Window)
		}
	}

	fTxt := canvas.NewText("Регестрация", color.White)
	fTxt.TextSize = 20
	fTxt.TextStyle.Bold = true
	fTxt.Alignment = fyne.TextAlignCenter

	label := widget.NewLabel("для того чтобы продолжить зарегестрируйтесь в системе")

	signButton := widget.NewButtonWithIcon("Войти", theme.AccountIcon(), func() {
		s.SignInWindow()
	})
	signButton.Enable()

	s.Window.SetContent(
		container.NewBorder(nil, nil, nil, nil,
			container.NewCenter(container.NewVBox(fTxt, label, canvas.NewLine(color.White), form, widget.NewActivity(), signButton))))
}

func (s *Settings) MainWindow() {
	var fromEntry = widget.NewEntry()
	fromEntry.SetPlaceHolder("Откуда")

	var toEntry = widget.NewEntry()
	toEntry.SetPlaceHolder("Куда")

	var dateButtonFrom = s.DateButton()
	dateButtonFrom.Text = "Туда (дата)"

	var dateButtonTo = s.DateButton()
	dateButtonTo.Text = "Обратно (дата)"

	grid := container.NewGridWithColumns(4, fromEntry, toEntry, dateButtonFrom, dateButtonTo)

	var buttonFind = widget.NewButtonWithIcon("Поиск", theme.SearchIcon(), func() {})

	var searchMenu = container.NewVBox(grid, widget.NewActivity(), buttonFind)
	buttonFind.OnTapped = func() {
		stack := container.NewVBox()
		ticket := models.NewTicketSearch() // инициализация структуры билета
		ticket.From = fromEntry.Text
		ticket.To = toEntry.Text
		ticket.Date_from = dateButtonFrom.Text
		ticket.Date_to = dateButtonTo.Text

		tickets, err := s.Storage.AllTickets()
		if err != nil {
			dialog.ShowInformation("Ошибка", err.Error(), s.Window)
		} else {
			fmt.Println(tickets, err)
			for i := 0; i < len(tickets); i++ {
				t := tickets[i]
				stack.Add(widget.NewButton(fmt.Sprintf("%s (%s) \u27F6 %s (%s) \n %s \u27F6 %s \n Цена: %.2f", t.DepartureFrom.City, t.DepartureFrom.Iata, t.ArrivalAt.City, t.ArrivalAt.Iata, t.DepartureTime, t.ArrivalTime, t.Price), func() {
					s.TicketWindow(&t)
				}))
			}
			searchMenu.Refresh()
			s.Window.SetContent(container.NewBorder(container.NewVBox(container.NewBorder(nil, nil, nil, s.ProfileButton()), container.NewCenter(searchMenu)), nil, nil, nil,  widget.NewCard("", "", container.NewVScroll(stack))))
		}
	}

	//grid2 := container.NewGridWrap(fyne.NewSize(50, 100), fromEntry, toEntry, dateButtonFrom, dateButtonTo)
	s.Window.SetContent(container.NewCenter(searchMenu))
}

func (s *Settings) TicketWindow(t *models.Ticket) {
	fmt.Println(t)
}



