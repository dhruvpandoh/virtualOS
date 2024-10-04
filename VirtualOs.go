package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	
)

var myApp fyne.App=app.New()

var myWindow fyne.Window = myApp.NewWindow("Virtual OS")


var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget

var img fyne.CanvasObject;
var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main(){

	img = canvas.NewImageFromFile("background.jpg");

	btn1=widget.NewButtonWithIcon("Weather App",theme.InfoIcon(), func(){
		showWeatherApp(myWindow);
	})

	btn2=widget.NewButtonWithIcon("Calculator App",theme.ContentAddIcon(), func(){
		showCalc();
	})

	btn3=widget.NewButtonWithIcon("Gallery App",theme.StorageIcon(), func(){
		showGallery();
	})

	btn4=widget.NewButtonWithIcon("TextEditor App",theme.DocumentIcon(), func(){
		showTextEditor();
	})


	DeskBtn=widget.NewButtonWithIcon("Home",theme.HomeIcon(), func(){
		myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
	})

	panelContent=container.NewVBox((container.NewGridWithColumns(5, DeskBtn, btn1, btn2, btn3, btn4)))

	myWindow.Resize(fyne.NewSize(1280, 720))
	myWindow.CenterOnScreen();

	myWindow.SetContent(
		container.NewBorder(panelContent, nil,nil,nil,img),
	)

	myWindow.ShowAndRun();
}