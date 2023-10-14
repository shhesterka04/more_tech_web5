package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Coordinates представляют широту и долготу
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// RouteResponse представляет ответ от OSRM API
type RouteResponse struct {
	Routes []struct {
		Geometry interface{} `json:"geometry"`
		Duration float64     `json:"duration"`
	} `json:"routes"`
}

// GeoJSON структура для линейного маршрута
type GeoJSON struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

// Result структура для вывода маршрута и времени пути
type Result struct {
	Path     GeoJSON `json:"path"`
	Duration float64 `json:"duration"` // время пути в секундах
}

// FetchRoute получает маршрут между стартом и финишем на основе типа транспорта и возвращает Result
func FetchRoute(start Coordinates, end Coordinates, transportType string) (Result, error) {
	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/%s/%f,%f;%f,%f?overview=full&geometries=geojson",
		transportType,
		start.Longitude, start.Latitude,
		end.Longitude, end.Latitude)

	resp, err := http.Get(url)
	if err != nil {
		return Result{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Result{}, err
	}

	var routeResponse RouteResponse
	err = json.Unmarshal(body, &routeResponse)
	if err != nil {
		return Result{}, err
	}

	var result Result
	if len(routeResponse.Routes) > 0 {
		geoJSONData, ok := routeResponse.Routes[0].Geometry.(map[string]interface{})
		if ok {
			result.Path.Type = geoJSONData["type"].(string)
			for _, coord := range geoJSONData["coordinates"].([]interface{}) {
				point := []float64{coord.([]interface{})[0].(float64), coord.([]interface{})[1].(float64)}
				result.Path.Coordinates = append(result.Path.Coordinates, point)
			}
			result.Duration = routeResponse.Routes[0].Duration
		}
		return result, nil
	}

	return Result{}, fmt.Errorf("No route found")
}

func main() {
	start := Coordinates{Latitude: 52.5200, Longitude: 13.4050} // Берлин
	end := Coordinates{Latitude: 48.8566, Longitude: 2.3522}    // Париж
	transportType := "driving"                                  // или "walking", "cycling" и т.д.

	result, err := FetchRoute(start, end, transportType)
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
