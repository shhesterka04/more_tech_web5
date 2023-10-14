package maps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchRoute(start Coordinates, end Coordinates, transportType string) ([][]float64, error) {
	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/%s/%f,%f;%f,%f?overview=full&geometries=geojson",
		transportType,
		start.Longitude, start.Latitude,
		end.Longitude, end.Latitude)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var routeResponse RouteResponse
	err = json.Unmarshal(body, &routeResponse)
	if err != nil {
		return nil, err
	}

	if len(routeResponse.Routes) > 0 {
		geoJSONData, ok := routeResponse.Routes[0].Geometry.(map[string]interface{})
		if ok {
			var coords [][]float64
			for _, coord := range geoJSONData["coordinates"].([]interface{}) {
				point := []float64{coord.([]interface{})[0].(float64), coord.([]interface{})[1].(float64)}
				coords = append(coords, point)
			}
			return coords, nil
		}
	}

	return nil, fmt.Errorf("No route found or unexpected format")
}
