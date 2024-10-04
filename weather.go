package main

import (
	// "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	// for http request
	"net/http"

)

func showWeatherApp(w fyne.Window){

	// a := app.New()
	// w := a.NewWindow("Hello")
	// w.Resize(fyne.NewSize(500,520))

	//usi weather api

	res, err:= http.Get("https://api.openweathermap.org/data/2.5/weather?q=jind&appid=5118a60eb8ab2b121b6450512dbf8922");

	if err !=nil{
		fmt.Print(err);
	}

	defer res.Body.Close();

	body , err:=ioutil.ReadAll(res.Body)
	if err !=nil{
		fmt.Print(err);
	}

	weather, err:= UnmarshalWelcome(body)
	if err !=nil{
		fmt.Print(err);
	}


	img:=canvas.NewImageFromFile("weatherPic.png");
	img.FillMode=canvas.ImageFillOriginal;

	lblDetails:=canvas.NewText("Weather Details",color.Black)
	lblDetails.TextStyle = fyne.TextStyle{Bold:true}

	lblCountry:=canvas.NewText(fmt.Sprintf("country %s", weather.Sys.Country),color.Black)
	lblWindSpeed:=canvas.NewText(fmt.Sprintf("wind speed %.2f", weather.Wind.Speed),color.Black)
	lblTemp:=canvas.NewText(fmt.Sprintf("Temperature %.2f", weather.Main.Temp),color.Black)
	lblHumidity:=canvas.NewText(fmt.Sprintln("Humidity ", weather.Main.Humidity),color.Black)



weatherContainer:=container.NewVBox(
			lblDetails,
			img,
			lblCountry,
			lblWindSpeed,
			lblTemp,
			lblHumidity,
			container.NewGridWithColumns(1,
				
			),	
		)


	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,weatherContainer),)
	w.Show();
}




func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Coord      Coord     `json:"coord"`     
	Weather    []Weather `json:"weather"`   
	Base       string    `json:"base"`      
	Main       Main      `json:"main"`      
	Visibility int64     `json:"visibility"`
	Wind       Wind      `json:"wind"`      
	Clouds     Clouds    `json:"clouds"`    
	Dt         int64     `json:"dt"`        
	Sys        Sys       `json:"sys"`       
	Timezone   int64     `json:"timezone"`  
	ID         int64     `json:"id"`        
	Name       string    `json:"name"`      
	Cod        int64     `json:"cod"`       
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`      
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`  
	TempMax   float64 `json:"temp_max"`  
	Pressure  int64   `json:"pressure"`  
	Humidity  int64   `json:"humidity"`  
	SeaLevel  int64   `json:"sea_level"` 
	GrndLevel int64   `json:"grnd_level"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"` 
}

type Weather struct {
	ID          int64  `json:"id"`         
	Main        string `json:"main"`       
	Description string `json:"description"`
	Icon        string `json:"icon"`       
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`  
	Gust  float64 `json:"gust"` 
}
