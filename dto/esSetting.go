package dto

// ESSetting ...
type ESSetting struct {
	Settings Setting `json:"settings"`
	Mappings Mapping `json:"mappings"`
}

// Mapping ...
type Mapping struct {
	Properties Property `json:"properties"`
}

// Property ...
type Property struct {
	MySrcLoc  Geo `json:"mysrc"`
	MyDestLoc Geo `json:"mydest"`
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
