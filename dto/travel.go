package dto

type Travel struct {
	ID                   string  `json:"id"`
	UserName             string  `json:"username"`
	Note                 string  `json:"note"`
	Mode                 string  `json:"mode"`
	SourceAddress        string  `json:"src_address"`
	DestinationAddress   string  `json:"dest_address"`
	SourceLatitude       float64 `json:"src_lat"`
	SourceLongitude      float64 `json:"src_long"`
	DestinationLatitude  float64 `json:"dest_lat"`
	DestinationLongitude float64 `json:"dest_long"`
	CreatedAt            string  `json:"created_at"`
	Status               string  `json:"status"`
	CompletedAt          string  `json:"completed_at"`
}
