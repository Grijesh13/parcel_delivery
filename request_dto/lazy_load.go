package request_dto

type LazyLoad struct {
	From      int     `json:"from"`
	Many      int     `json:"many"`	// if not mentioned use default value
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
