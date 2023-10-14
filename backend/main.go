package main

import (
	"encoding/json"
	"fmt"
	"more_tech_web5/backend/maps"
)

//55.789268,37.540339 - старт

//55.781400,37.532215 - финиш

func main() {
	start := maps.Coordinates{Latitude: 37.540339, Longitude: 55.789268} // Берлин
	end := maps.Coordinates{Latitude: 37.532215, Longitude: 55.781400}   // Париж
	transportType := "car"                                               // или "walking", "cycling" и т.д.

	result, err := maps.FetchRoute(start, end, transportType)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(jsonResult))
}
