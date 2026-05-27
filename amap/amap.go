package amap

import (
	"encoding/xml"

	"github.com/codingsince1985/geo-golang"
)

type (
	baseURL         string
	geocodeResponse struct {
		XMLName  xml.Name `xml:"response"`
		Status   int      `xml:"status"`
		Info     string   `xml:"info"`
		Infocode int      `xml:"infocode"`
		Count    int      `xml:"count"`
		Geocodes []struct {
			FormattedAddress string `xml:"formatted_address"`
			Country          string `xml:"country"`
			Province         string `xml:"province"`
			Citycode         string `xml:"citycode"`
			City             string `xml:"city"`
			District         string `xml:"district"`
			Adcode           string `xml:"adcode"`
			Street           string `xml:"street"`
			Number           string `xml:"number"`
			Location         string `xml:"location"`
			Level            string `xml:"level"`
		} `xml:"geocodes>geocode"`
		Regeocode struct {
			FormattedAddress string `xml:"formatted_address"`
			AddressComponent struct {
				Country      string `xml:"country"`
				Township     string `xml:"township"`
				District     string `xml:"district"`
				Adcode       string `xml:"adcode"`
				Province     string `xml:"province"`
				Citycode     string `xml:"citycode"`
				StreetNumber struct {
					Number    string `xml:"number"`
					Location  string `xml:"location"`
					Direction string `xml:"direction"`
					Distance  string `xml:"distance"`
					Street    string `xml:"street"`
				} `xml:"streetNumber"`
			} `xml:"addressComponent"`
		} `xml:"regeocode"`
	}
)

const (
	statusOK = 1
)

var r = 1000

// Geocoder constructs AMAP geocoder
func Geocoder(key string, radius int, baseURLs ...string) geo.Geocoder {
	_ = "STUB: not implemented"
	return *new(geo.Geocoder)
}

func getURL(apiKey string, baseURLs ...string) string { _ = "STUB: not implemented"; return "" }

// GeocodeURL https://restapi.amap.com/v3/geocode/geo?&output=XML&key=APPKEY&address=ADDRESS
func (b baseURL) GeocodeURL(address string) string { _ = "STUB: not implemented"; return "" }

// ReverseGeocodeURL https://restapi.amap.com/v3/geocode/regeo?output=XML&key=APPKEY&radius=1000&extensions=all&location=31.225696563611,121.49884033194
func (b baseURL) ReverseGeocodeURL(l geo.Location) string { _ = "STUB: not implemented"; return "" }

func (r *geocodeResponse) Location() (*geo.Location, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *geocodeResponse) Address() (*geo.Address, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseAmapResult(r *geocodeResponse) *geo.Address { _ = "STUB: not implemented"; return nil }
