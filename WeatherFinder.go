package main

import (
	"net/http"
	"html/template"
	"fmt"
	"encoding/json"
)

type weatherReport struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		TempMin   float64 `json:"temp_min"`
		TempMax   float64 `json:"temp_max"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}


func main(){

	weatherFinder := func(w http.ResponseWriter, r *http.Request){

		endpoint := "http://api.openweathermap.org/data/2.5/weather?q="
	
		appId := "3e8050701dd730c369eba684d26b4e4f"  
		
		location := r.FormValue("city")

		req := endpoint + location + "&APPID=" + appId

		fmt.Println(location)

		resp,er := http.Get(req)

		if er != nil{
			print(er)
		}
	
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)

		var Data weatherReport
		
		err := decoder.Decode(&Data)

		if err == nil{
			print(err)
		}
		
		wr,er := template.ParseFiles("weather.html")

		if er != nil{
			print("Something went wrong")
		}else{
			wr.Execute(w,Data)
		}
	
		fmt.Println(wr)
	}

	Iweather := func(w http.ResponseWriter, r *http.Request){
		t,err := template.ParseFiles("getCity.html")
		if err != nil{
			fmt.Println("panic")
		}
		data := "fine"
		t.Execute(w,data)
	}

	http.HandleFunc("/",Iweather)

	http.HandleFunc("/FindWeather",weatherFinder)

	http.ListenAndServe(":9005",nil)
	
} 
