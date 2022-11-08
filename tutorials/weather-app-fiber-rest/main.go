// source How to Consume/GET Weather Data Using Golang FIBER /HTTP REST API - PART 45(a)
// https://www.youtube.com/watch?v=bdU2NJF7XT0
// go mod init weatherapp
// go get github.com/gofiber/fiber/v2
// go get github.com/joho/godotenv

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const CITY = "Dallas"

func GetEnv(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

// weather struct
type WeatherResponse struct {
	// Name string `json:"name"`
	// Main Main   `json:"main"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	//Current_Weather Current_Weather `json:"current_weather"`
	Current_Weather Current_Weather `json:"current_weather"`
}

type Current_Weather struct {
	Temperature float64 `json:"temperature"`
	// Windspeed     float64 `json:"windspeed"`
	// WindDirection float64 `json:"winddirection"`
	// WeatherCode   int     `json:"weathercode"`
	// Time          string  `json:"time"`
}

// type Main struct {
// 	// Temp      float64 `json:"temp"`
// 	// FeelsLike float64 `json:"feels_like"`
// 	// TempMin   float64 `json:"temp_min"`
// 	// TempMax   float64 `json:"temp_max"`
// 	// Pressure  float64 `json:"pressure"`
// 	// Humidity  float64 `json:"humidity"`
// 	Latitude float64 `json:"latitude"`
// }

func main() {
	// println("hi")
	// LoadEnv()
	// APIKEY := GetEnv("APIKEY")
	// fmt.Println("apikey is not used", APIKEY)
	//URL := "https://api.openweathermap.org/data/2.5/weather?q=" + CITY + "&appid=" + APIKEY + "&units=metric"
	URL := "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&hourly=temperature_2m"
	//URL := "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current_weather=true"
	response, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)
	//ioutil package
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		println(err)
	}
	//fmt.Println(jsonBytes)
	defer func() {
		e := response.Body.Close()
		if e != nil {
			log.Fatal(e)
		}
	}()

	var weatherResponse WeatherResponse
	er := json.Unmarshal(jsonBytes, &weatherResponse)
	if er != nil {
		log.Fatal(er)
	}
	// //
	fmt.Printf("%+v\n", weatherResponse)
}
