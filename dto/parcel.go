package dto

type Parcel struct {
	ID                   string  `json:"id"`
	UserName             string  `json:"username"`
	Note                 string  `json:"note"`
	Length               int     `json:"length"`
	Breadth              int     `json:"breadth"`
	Height               int     `json:"height"`
	Weight               int     `json:"weight"`
	Category             string  `json:"category"`
	SourceAddress        string  `json:"src_address"`
	DestinationAddress   string  `json:"src_address"`
	SourceLatitude       float64 `json:"src_lat"`
	SourceLongitude      float64 `json:"src_long"`
	DestinationLatitude  float64 `json:"dest_lat"`
	DestinationLongitude float64 `json:"dest_long"`
	CreatedAt            string  `json:"created_at"`
}
