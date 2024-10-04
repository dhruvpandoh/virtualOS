package main

import (
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"io/ioutil"
    "log"
	"fmt"
	"strings"
	"strconv"

)

func showGallery() {
	// a := app.New()
	// w := a.NewWindow("Gallery")
	// w.Resize(fyne.NewSize(780,520))
	root_src:= "E:\\goProject\\img";

	files, err := ioutil.ReadDir(root_src)
    if err != nil {
        log.Fatal(err)
    }

	tabs:= container.NewAppTabs()
	var indx int=1;
    for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())

		if file.IsDir()==false{
			extension:= strings.Split(file.Name(), ".")[1];

			if extension =="png" || extension=="jpg"{
				//  ArrayOfPics = append(ArrayOfPics,root_src+"\\"+file.Name());
				 image:=canvas.NewImageFromFile(root_src+"\\"+file.Name())
				 image.FillMode = canvas.ImageFillOriginal
				 tabs.Append(container.NewTabItem("Image"+strconv.Itoa(indx),image));
				 indx++;
			}
		}
    }

	tabs.SetTabLocation(container.TabLocationLeading)
	// w.SetContent(tabs);
	// w.ShowAndRun()

	w:=myApp.NewWindow("myGallery")	
	w.Resize(fyne.NewSize(500,280));

	w.SetContent(
		container.NewBorder(DeskBtn,nil,nil,nil,tabs),
	)

	w.Show();
}