package model

// Request ...
type Request struct {
	Date   string `json:"date"`
	Views  string `json:"views,omitempty"`
	Clicks string `json:"clicks,omitempty"`
	Cost   string `json:"cost,omitempty"`
}
