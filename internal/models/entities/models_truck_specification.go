package entities

type Truck struct {
	Dimensions        Dimensions
	Weight            int   // Total weight including cargo
	AxleWeight        []int // Weight per axle
	Hazmat            bool
	CargoType         string
	Refrigeration     bool
	FuelCapacity      int
	FuelStations      []FuelStation // List of preferred fuel stations
	DriverHours       int
	RestAreas         []string
	RouteRestrictions []string
	TollSystem        string
	TrafficPatterns   map[string]string // key: location, value: traffic pattern
	TurnRadius        int
	LoadZones         []string
	WeatherPatterns   map[string]string // key: location, value: weather pattern
	RoadGrades        map[string]string // key: location, value: road grade
	BridgeHeights     map[string]int
	EmergencyServices []string
	Insurance         map[string]string // key: insurance type, value: details
	Borders           []string
}
