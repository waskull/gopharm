package domain

// Date represents a date with day, month, and year.
type Date struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year"`
}

// TimeRange represents a time range with start and end times. Useful
// for specifying the time range of games.
type TimeRange struct {
	Start TimeOfDay `json:"start"`
	End   TimeOfDay `json:"end"`
}

// TimeOfDay represents a time of day with hour and minute.
type TimeOfDay struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}
