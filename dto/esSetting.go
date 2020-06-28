package dto

type ESSetting struct {
	Settings Setting `json:"settings"`
	Mappings Mapping `json:"mappings"`
}

type Mapping struct {
	Properties Property `json:"properties"`
}

type Property struct {
	MySrcLoc   Geo `json:"mysrc"`
	MyDestLoc  Geo `json:"mydest"`
}

type Geo struct {
	Type string `json:"type"`
}

type Setting struct {
	Shards   int `json:"number_of_shards"`
	Replicas int `json:"number_of_replicas"`
}
