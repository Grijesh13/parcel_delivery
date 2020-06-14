package dto

type ESParcel struct {
	MyLoc                Loc    `json:"myloc"`
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
	Status               string  `json:"status"`
	Price                int     `json:"price"`
	CompletedAt          string  `json:"completed_at"`
}

type Loc struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lon"`
}
