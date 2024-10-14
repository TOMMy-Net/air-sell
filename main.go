package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/TOMMy-Net/air-sell/internal"
)

func main() {
	myApp := app.New()
	//myApp.Settings().SetTheme(theme.DarkTheme())

	mainWindow := myApp.NewWindow("Air Buy")
	mainWindow.SetIcon(internal.GetIcon())
	mainWindow.Resize(fyne.Size{Width: 1000, Height: 700})
	mainWindow.SetMaster()
	mainWindow.CenterOnScreen()
	internal.SignInWindow(mainWindow)
	mainWindow.Show()
	myApp.Run()

}
