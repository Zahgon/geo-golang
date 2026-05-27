// Package google is a geo-golang based Google Geo Location API
// https://developers.google.com/maps/documentation/geocoding/intro
package google

import (
	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Results []struct {
			FormattedAddress  string                   `json:"formatted_address"`
			AddressComponents []googleAddressComponent `json:"address_components"`
			Geometry          struct {
				Location geo.Location
			}
		}
		Status string `json:"status"`
	}
	googleAddressComponent struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	}
)

const (
	statusOK                   = "OK"
	statusNoResults            = "ZERO_RESULTS"
	componentTypeHouseNumber   = "street_number"
	componentTypeStreetName    = "route"
	componentTypeSuburb        = "sublocality"
	componentTypeLocality      = "locality"
	componentTypeStateDistrict = "administrative_area_level_2"
	componentTypeState         = "administrative_area_level_1"
	componentTypeCountry       = "country"
	componentTypePostcode      = "postal_code"
)

// Geocoder constructs Google geocoder
func Geocoder(apiKey string, baseURLs ...string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

func getURL(apiKey string, baseURLs ...string) string { _ = "STUB: not implemented"; return "" }

func (b baseURL) GeocodeURL(address string) string { _ = "STUB: not implemented"; return "" }

func (b baseURL) ReverseGeocodeURL(l geo.Location) string { _ = "STUB: not implemented"; return "" }

func (r *geocodeResponse) Location() (*geo.Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseGoogleResult(r *geocodeResponse) *geo.Address { _ = "STUB: not implemented"; return nil }
