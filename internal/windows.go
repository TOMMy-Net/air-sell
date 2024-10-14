package internal

import (
	"image/color"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"

	"fyne.io/fyne/v2/widget"
)

func SignInWindow(w fyne.Window) {

	form := widget.NewForm(
		widget.NewFormItem("email", widget.NewEntry()),
		widget.NewFormItem("password", widget.NewPasswordEntry()),
	)

	form.SubmitText = "Войти"
	form.Orientation = widget.Vertical

	form.OnSubmit = func() {

	}

	label := widget.NewLabel("для того чтобы продолжить войдите в систему")

	fTxt := canvas.NewText("Добро Пожаловать", color.White)
	fTxt.TextSize = 20
	fTxt.TextStyle.Bold = true
	fTxt.Alignment = fyne.TextAlignCenter

	regButton := widget.NewButtonWithIcon("Зарегестрироваться", theme.AccountIcon(), func() {
		SignUpWindow(w)
	})
	regButton.Enable()

	w.SetContent(
		container.NewBorder(nil, nil, nil, nil,
			container.NewCenter(container.NewVBox(fTxt, label, form, widget.NewActivity(), regButton))))

}

func SignUpWindow(w fyne.Window) {
	form := widget.NewForm(
		widget.NewFormItem("email", widget.NewEntry()),
		widget.NewFormItem("password", widget.NewPasswordEntry()),
		widget.NewFormItem("repeat the password", widget.NewPasswordEntry()),
	)

	form.SubmitText = "Зарегестрироваться"
	form.Orientation = widget.Vertical

	form.OnSubmit = func() {

	}

	fTxt := canvas.NewText("Регестрация", color.White)
	fTxt.TextSize = 20
	fTxt.TextStyle.Bold = true
	fTxt.Alignment = fyne.TextAlignCenter

	label := widget.NewLabel("для того чтобы продолжить зарегестрируйтесь в системе")

	signButton := widget.NewButtonWithIcon("Войти", theme.AccountIcon(), func() {
		SignInWindow(w)
	})
	signButton.Enable()

	w.SetContent(
		container.NewBorder(nil, nil, nil, nil,
			container.NewCenter(container.NewVBox(fTxt,label, form, widget.NewActivity(), signButton))))
}
