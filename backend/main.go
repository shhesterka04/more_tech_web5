package main

import (
	"encoding/json"
	"fmt"
	"more_tech_web5/backend/maps"
)

func main() {
	start := maps.Coordinates{Latitude: 52.5200, Longitude: 13.4050} // Берлин
	end := maps.Coordinates{Latitude: 48.8566, Longitude: 2.3522}    // Париж
	transportType := "driving"                                       // или "walking", "cycling" и т.д.

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
