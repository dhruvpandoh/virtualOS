package main

import (
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"fyne.io/fyne/v2/dialog"
	"io/ioutil"
	"fyne.io/fyne/v2/storage"
    // "log"
	// "fmt"
	// "strings"

)

var fileIndex int=1;

func showTextEditor() {
	// a := app.New()
	w := myApp.NewWindow("Your_Editor")
	w.Resize(fyne.NewSize(780,520))

	content:=container.NewVBox(

		container.NewHBox(

			widget.NewLabel("Your_Editor"),
		),
	)

	content.Add(widget.NewButton("Add New File", func(){

		content.Add(widget.NewLabel("New File"+strconv.Itoa(fileIndex)))
		fileIndex++;
	}))

	input:=widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter text....")	
	input.Resize(fyne.NewSize(400, 400))

	saveBtn:=widget.NewButton("Save your File", func(){

		saveFileDialog:= dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error){

				textData:=[]byte(input.Text)

				uc.Write(textData);
			},w)

			saveFileDialog.SetFileName("New File"+strconv.Itoa(fileIndex-1)+".txt")
			saveFileDialog.Show()
	})


	openBtn:=widget.NewButton("open Text File", func(){

		openFileDialog:=dialog.NewFileOpen(

			func(r fyne.URIReadCloser, _ error){

				ReadData,_:=ioutil.ReadAll(r)

				output:= fyne.NewStaticResource("New File", ReadData);

				viewData:=widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))

				// to open new window on current App
				w:=fyne.CurrentApp().NewWindow(

					string(output.StaticName))

					w.SetContent(container.NewScroll(viewData))
					w.Resize(fyne.NewSize(400,400))
					w.Show();
			},w)

				openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))

				openFileDialog.Show();
	})

	//w.SetContent(
	editors:=container.NewVBox(
			content,
			input,

			container.NewHBox(
				saveBtn,
				openBtn,
			),
		)//,
	// )

	w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,editors),
	)
	w.Show();
	// w.ShowAndRun()
}