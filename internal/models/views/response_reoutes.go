package views

type RouteSummary struct {
	LengthInMeters        int
	TravelTimeInSeconds   int
	TrafficDelayInSeconds int
}

type Route struct {
	Summary RouteSummary `json:"summary"`
}

type RouteResponse struct {
	Routes []Route `json:"routes"`
}
