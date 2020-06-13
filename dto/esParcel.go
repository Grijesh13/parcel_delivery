package dto

type ESParcel struct {
	MyLoc Loc    `json:"myloc"`
	Info  string `json:"info"`
}

type Loc struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"lon"`
}
