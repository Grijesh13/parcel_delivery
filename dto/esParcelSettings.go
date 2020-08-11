package dto

// ESParcelSetting ...
type ESParcelSetting struct {
	Settings Setting `json:"settings"`
	Mappings ParcelMapping `json:"mappings"`
}

// ParcelMapping ...
type ParcelMapping struct {
	Properties ParcelProperty `json:"properties"`
}

// ParcelProperty ...
type ParcelProperty struct {
	MySrcLoc  Geo `json:"mysrc"`
	MyDestLoc Geo `json:"mydest"`
	PickUpStart ESDate `json:"pick_up_start"`
	PickUpEnd ESDate `json:"pick_up_end"`
}
