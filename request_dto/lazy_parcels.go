package request_dto

type LazyParcels struct {
	From      int     `json:"from"`
	Many      int     `json:"many"`	// if not mentioned use default value
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
