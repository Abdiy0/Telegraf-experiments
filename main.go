package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const APIKey = "6cfd147d4ff64062bfc20745210206"

const WeatherAPIURL = "https://api.weatherapi.com/v1/current.json"

const CityName = "Toronto"

type Location struct {
	Country  string  `json:"country"`
	Timezone string  `json:"tz_id"`
	Weather  float64 `json:"temp_c"`
}

type Response struct {
	Location Location `json:"location"`
}

func main() {

	url := fmt.Sprintf("%s?key=%s&q=%s", WeatherAPIURL, APIKey, CityName)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println(string(body))

	// var result map[string]interface{}
	var result Response
	json.NewDecoder(resp.Body).Decode(&result)
	// log.Println(result["location"]["country"])
	log.Println(result)
	log.Printf("%#v\n", result)

	b, err := json.MarshalIndent(result, "", " ")
	if err != nil {

		log.Println("err")

	}
	log.Println(string(b))

}
