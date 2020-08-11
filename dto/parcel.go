package dto

// Parcel ...
type Parcel struct {
	ID                   string  `json:"id"`
	UserName             string  `json:"username"`
	Note                 string  `json:"note"`
	Items                []Item  `json:"items"`
	SQLItems             string  `json:"sql_items"`
	SourceAddress        string  `json:"src_address"`
	DestinationAddress   string  `json:"dest_address"`
	SourceLatitude       float64 `json:"src_lat"`
	SourceLongitude      float64 `json:"src_long"`
	DestinationLatitude  float64 `json:"dest_lat"`
	DestinationLongitude float64 `json:"dest_long"`
	PickUpStart          string  `json:"pick_up_start"`
	PickUpEnd            string  `json:"pick_up_end"`
	CreatedAt            string  `json:"created_at"`
	Status               string  `json:"status"`
	Price                int     `json:"price"`
	IsNegotiable         bool    `json:"is_negotiable"`
	CompletedAt          string  `json:"completed_at"`
}

// Item ...
type Item struct {
	Number   int    `json:"number"`
	Length   int    `json:"length"`
	Breadth  int    `json:"breadth"`
	Height   int    `json:"height"`
	Weight   int    `json:"weight"`
	Category string `json:"category"`
}
