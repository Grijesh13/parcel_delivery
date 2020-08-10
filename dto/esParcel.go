package dto

// ESParcel ...
type ESParcel struct {
	MySrcLoc             Loc      `json:"mysrc"`
	MyDestLoc            Loc      `json:"mydest"`
	PickUpStart          string   `json:"pick_up_start"`
	UserName             string   `json:"username"`
	Note                 string   `json:"note"`
	Items                []Item   `json:"items"`
	SourceAddress        string   `json:"src_address"`
	DestinationAddress   string   `json:"dest_address"`
	SourceLatitude       float64  `json:"src_lat"`
	SourceLongitude      float64  `json:"src_long"`
	DestinationLatitude  float64  `json:"dest_lat"`
	DestinationLongitude float64  `json:"dest_long"`
	CreatedAt            string   `json:"created_at"`
	Status               string   `json:"status"`
	Price                int      `json:"price"`
	CompletedAt          string   `json:"completed_at"`
	IsNegotiable         bool     `json:"is_negotiable"`
	ShipDate             string   `json:"ship_date"`
	NumberItems          int      `json:"num_items"`
	NetWeight            int      `json:"net_weight"`
	Categories           []string `json:"categories"`
}

// Loc ...
type Loc struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lon"`
}
