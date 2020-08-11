package dto

// ESTravelSetting ...
type ESTravelSetting struct {
	Settings Setting `json:"settings"`
	Mappings TravelMapping `json:"mappings"`
}

// TravelMapping ...
type TravelMapping struct {
	Properties TravelProperty `json:"properties"`
}

// TravelProperty ...
type TravelProperty struct {
	MySrcLoc  Geo `json:"mysrc"`
	MyDestLoc Geo `json:"mydest"`
	StartDate ESDate `json:"start_date"`
	EndDate ESDate `json:"end_date"`
}

type ESDate struct {
	Type   string `json:"type"`
	Format string `json:"format"`
}

// Geo ...
type Geo struct {
	Type string `json:"type"`
}

// Setting ...
type Setting struct {
	Shards   int `json:"number_of_shards"`
	Replicas int `json:"number_of_replicas"`
}
