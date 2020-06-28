package request_dto

type LazyLoad struct {
	From          int     `json:"from"`
	Many          int     `json:"many"`	// if not mentioned use default value
	SrcLatitude   float64 `json:"src_latitude"`
	SrcLongitude  float64 `json:"src_longitude"`
	SrcDistance   int     `json:"src_distance"`
	DestLatitude  float64 `json:"dest_latitude"`
	DestLongitude float64 `json:"dest_longitude"`
	DestDistance  int     `json:"dest_distance"`
	DestGiven     bool    `json:"dest_given"`
}
