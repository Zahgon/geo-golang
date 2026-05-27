// Package osm provides common types for OpenStreetMap used by various providers
// and some helper functions to reduce code repetition across specific client implementations.
package osm

// Address contains address fields specific to OpenStreetMap
type Address struct {
	HouseNumber   string `json:"house_number"`
	Road          string `json:"road"`
	Pedestrian    string `json:"pedestrian"`
	Footway       string `json:"footway"`
	Cycleway      string `json:"cycleway"`
	Highway       string `json:"highway"`
	Path          string `json:"path"`
	Suburb        string `json:"suburb"`
	City          string `json:"city"`
	Town          string `json:"town"`
	Village       string `json:"village"`
	Hamlet        string `json:"hamlet"`
	County        string `json:"county"`
	Country       string `json:"country"`
	CountryCode   string `json:"country_code"`
	State         string `json:"state"`
	StateDistrict string `json:"state_district"`
	Postcode      string `json:"postcode"`
}

// Locality checks different fields for the locality name
func (a Address) Locality() string { _ = "STUB: not implemented"; return "" }

// Street checks different fields for the street name
func (a Address) Street() string { _ = "STUB: not implemented"; return "" }
