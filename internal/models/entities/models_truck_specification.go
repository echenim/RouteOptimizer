package entities

type Location struct {
	Address   string  `json:"address"` // Using an address string for simplicity. Consider using more complex types like GPS coordinates.
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Truck struct {
	Dimensions        Dimensions      `json:"dimensions"`
	Weight            int             `json:"weight"`
	AxleWeight        []int           `json:"axle_weight"`
	Hazmat            bool            `json:"hazmat"`
	CargoType         CargoType       `json:"cargo_type"`
	Refrigeration     bool            `json:"refrigeration"`
	FuelCapacity      int             `json:"fuel_capacity"`
	FuelStations      []FuelStation   `json:"fuel_stations"`
	DriverHours       int             `json:"driver_hours"`
	RestAreas         []string        `json:"rest_areas"`
	RouteRestrictions []string        `json:"route_restrictions"`
	TollSystem        string          `json:"toll_system"`
	TrafficPatterns   TrafficPatterns `json:"traffic_patterns"`
	TurnRadius        int             `json:"turn_radius"`
	LoadZones         []string        `json:"load_zones"`
	WeatherPatterns   WeatherPatterns `json:"weather_patterns"`
	RoadGrades        RoadGrades      `json:"road_grades"`
	BridgeHeights     BridgeHeights   `json:"bridge_heights"`
	EmergencyServices []string        `json:"emergency_services"`
	Insurance         Insurance       `json:"insurance"`
	Borders           []string        `json:"borders"`
	CurrentLocation   Location        `json:"current_location"`
	Destination       Location        `json:"destination"`
}

// Enums or structured types for better control and clarity.
type CargoType string

const (
	CargoTypeGeneral    CargoType = "General"
	CargoTypePerishable CargoType = "Perishable"
	CargoTypeHazardous  CargoType = "Hazardous"
)

type (
	TrafficPatterns map[string]TrafficPattern
	TrafficPattern  string
)

const (
	TrafficPatternHigh   TrafficPattern = "High"
	TrafficPatternMedium TrafficPattern = "Medium"
	TrafficPatternLow    TrafficPattern = "Low"
)

type (
	WeatherPatterns map[string]WeatherPattern
	WeatherPattern  string
)

const (
	WeatherPatternSunny WeatherPattern = "Sunny"
	WeatherPatternRainy WeatherPattern = "Rainy"
	WeatherPatternSnowy WeatherPattern = "Snowy"
)

type (
	RoadGrades map[string]RoadGrade
	RoadGrade  string
)

const (
	RoadGradeFlat  RoadGrade = "Flat"
	RoadGradeHilly RoadGrade = "Hilly"
	RoadGradeSteep RoadGrade = "Steep"
)

type BridgeHeights map[string]int

type Insurance map[string]string
