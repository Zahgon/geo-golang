// Package yandex is a geo-golang based Yandex Maps Location API
package yandex

import (
	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		Response struct {
			GeoObjectCollection struct {
				MetaDataProperty struct {
					GeocoderResponseMetaData struct {
						Request string `json:"request"`
						Found   string `json:"found"`
						Results string `json:"results"`
					} `json:"GeocoderResponseMetaData"`
				} `json:"metaDataProperty"`
				FeatureMember []*yandexFeatureMember `json:"featureMember"`
			} `json:"GeoObjectCollection"`
		} `json:"response"`
	}

	yandexFeatureMember struct {
		GeoObject struct {
			MetaDataProperty struct {
				GeocoderMetaData struct {
					Kind      string `json:"kind"`
					Text      string `json:"text"`
					Precision string `json:"precision"`
					Address   struct {
						CountryCode string `json:"country_code"`
						PostalCode  string `json:"postal_code"`
						Formatted   string `json:"formatted"`
						Components  []struct {
							Kind string `json:"kind"`
							Name string `json:"name"`
						} `json:"Components"`
					} `json:"Address"`
				} `json:"GeocoderMetaData"`
			} `json:"metaDataProperty"`
			Description string `json:"description"`
			Name        string `json:"name"`
			BoundedBy   struct {
				Envelope struct {
					LowerCorner string `json:"lowerCorner"`
					UpperCorner string `json:"upperCorner"`
				} `json:"Envelope"`
			} `json:"boundedBy"`
			Point struct {
				Pos string `json:"pos"`
			} `json:"Point"`
		} `json:"GeoObject"`
	}
)

const (
	componentTypeHouseNumber   = "house"
	componentTypeStreetName    = "street"
	componentTypeLocality      = "locality"
	componentTypeStateDistrict = "area"
	componentTypeState         = "province"
	componentTypeCountry       = "country"
)

// Geocoder constructs Yandex geocoder
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

// Yandex return geo coord in format "long lat"

func (r *geocodeResponse) Address() (*geo.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseYandexResult(r *yandexFeatureMember) *geo.Address { _ = "STUB: not implemented"; return nil }
