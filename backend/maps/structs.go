package maps

type RouteResponse struct {
	Routes []struct {
		Geometry interface{} `json:"geometry"`
		Duration float64     `json:"duration"`
	} `json:"routes"`
}

type GeoJSON struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type Result struct {
	Path     GeoJSON `json:"path"`
	Duration float64 `json:"duration"`
}
type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type MapRoute struct {
	Start         Coordinates `json:"start"`
	End           Coordinates `json:"end"`
	TransportType string      `json:"transportType"`
}
