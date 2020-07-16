package dto

// ESTravel ...
type ESTravel struct {
	MySrcLoc             Loc     `json:"mysrc"`
	MyDestLoc            Loc     `json:"mydest"`
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
	StartDate            string  `json:"start_date"`
	EndDate              string  `json:"end_date"`
	CompletedAt          string  `json:"completed_at"`
}
