package main

import (
	"encoding/json"
	"fmt"
	"more_tech_web5/backend/maps"
)

//55.814396, 38.978727 - старт

//55.813799, 38.976005 - финиш

func main() {
	start := maps.Coordinates{Latitude: 55.814396, Longitude: 38.978727}
	end := maps.Coordinates{Latitude: 55.813799, Longitude: 38.976005}
	transportType := "car" // или "foot", "bike" и т.д.

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
