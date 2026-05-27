package main

import (
	"github.com/codingsince1985/geo-golang"
)

const (
	addr             = "Melbourne VIC"
	lat, lng         = -37.813611, 144.963056
	radius           = 50
	zoom             = 18
	addrFR           = "Champs de Mars Paris"
	addrAmap         = "北京市海淀区清河街道西二旗西路领秀新硅谷" // China only
	latFR, lngFR     = 48.854395, 2.304770
	latAmap, lngAmap = 40.05510, 116.309866
)

func main() {
	ExampleGeocoder()
}

// ExampleGeocoder demonstrates the different geocoding services
func ExampleGeocoder() { _ = "STUB: not implemented"; return }

// To use only with french locations or addresses,
// or values ​​could be estimated and will be false

// Chained geocoder will fallback to subsequent geocoders

// Output: Google Geocoding API
// Melbourne VIC location is (-37.813611, 144.963056)
// Address of (-37.813611,144.963056) is 197 Elizabeth St, Melbourne VIC 3000, Australia
// Detailed address: &geo.Address{FormattedAddress:"197 Elizabeth St, Melbourne VIC 3000, Australia",
// 	Street:"Elizabeth Street", HouseNumber:"197", Suburb:"", Postcode:"3000", State:"Victoria",
// 	StateDistrict:"Melbourne City", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
//
// Mapquest Nominatim
// Melbourne VIC location is (-37.814218, 144.963161)
// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Melbourne, City of Melbourne,
// 	Greater Melbourne, Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne",
// 	Postcode:"3000", State:"Victoria", StateDistrict:"", County:"City of Melbourne", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
//
// Mapquest Open streetmaps
// Melbourne VIC location is (-37.814218, 144.963161)
// Address of (-37.813611,144.963056) is Elizabeth Street, Melbourne, Victoria, AU
// Detailed address: &geo.Address{FormattedAddress:"Elizabeth Street, 3000, Melbourne, Victoria, AU",
// 	Street:"Elizabeth Street", HouseNumber:"", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"",
// 	County:"", Country:"", CountryCode:"AU", City:"Melbourne"}
//
// OpenCage Data
// Melbourne VIC location is (-37.814217, 144.963161)
// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Melbourne VIC 3000, Australia
// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Melbourne VIC 3000, Australia",
// 	Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne (3000)", Postcode:"3000", State:"Victoria",
// 	StateDistrict:"", County:"City of Melbourne", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
//
// HERE Geocoder API
// Melbourne VIC location is (-37.817530, 144.967150)
// Address of (-37.813611,144.963056) is 197 Elizabeth St, Melbourne VIC 3000, Australia
// Detailed address: &geo.Address{FormattedAddress:"197 Elizabeth St, Melbourne VIC 3000, Australia", Street:"Elizabeth St",
// 	HouseNumber:"197", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"", Country:"Australia",
// 	CountryCode:"AUS", City:"Melbourne"}
//
// Bing Geocoding API
// Melbourne VIC location is (-37.824299, 144.977997)
// Address of (-37.813611,144.963056) is Elizabeth St, Melbourne, VIC 3000
// Detailed address: &geo.Address{FormattedAddress:"Elizabeth St, Melbourne, VIC 3000", Street:"Elizabeth St",
// 	HouseNumber:"", Suburb:"", Postcode:"3000", State:"", StateDistrict:"", County:"", Country:"Australia", CountryCode:"", City:"Melbourne"}
//
// Baidu Geocoding API
// Melbourne VIC location is (31.227015, 121.456967)
// Address of (-37.813611,144.963056) is 341 Little Bourke Street, Melbourne, Victoria, Australia
// Detailed address: &geo.Address{FormattedAddress:"341 Little Bourke Street, Melbourne, Victoria, Australia", Street:"Little Bourke Street", HouseNumber:"341", Suburb:"", Postcode:"", State:"Victoria", StateCode:"", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AUS", City:"Melbourne"}
//
// Amap Geocoding API
// 北京市海淀区清河街道西二旗西路领秀新硅谷 location is (40.05510, 116.309866)
// Address of (40.05510, 116.309866) is 北京市海淀区清河街道领秀新硅谷领秀新硅谷B区
// Detailed address: &geo.Address{FormattedAddress:"北京市海淀区清河街道领秀新硅谷领秀新硅谷B区", Street:"西二旗西路", HouseNumber:"2号楼", Suburb:"海 淀区", Postcode:"", State:"北京市", StateCode:"", StateDistrict:"", County:"", Country:"中国", CountryCode:"", City:""}
//
// Mapbox API
// Melbourne VIC location is (-37.814200, 144.963200)
// Address of (-37.813611,144.963056) is Elwood Park Playground, Melbourne, Victoria 3000, Australia
// Detailed address: &geo.Address{FormattedAddress:"Elwood Park Playground, Melbourne, Victoria 3000, Australia",
// 	Street:"Elwood Park Playground", HouseNumber:"", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"",
// 	County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
//
// OpenStreetMap
// Melbourne VIC location is (-37.814217, 144.963161)
// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne,
// 	Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria",
// 	StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
//
// PickPoint
// Melbourne VIC location is (-37.814217, 144.963161)
// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne,
// 	Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria",
// 	StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
//
// LocationIQ
// Melbourne VIC location is (-37.814217, 144.963161)
// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne,
// 	Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria",
// 	StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}
//
// ArcGIS
// Melbourne VIC location is (-37.817530, 144.967150)
// Address of (-37.813611,144.963056) is Melbourne's Gpo
// Detailed address: &geo.Address{FormattedAddress:"Melbourne's Gpo", Street:"350 Bourke Street Mall", HouseNumber:"350", Suburb:"", Postcode:"3000", State:"Victoria", StateDistrict:"", County:"", Country:"", CountryCode:"AUS", City:""}
//
// geocod.io
// Melbourne VIC location is (28.079357, -80.623618)
// got <nil> address
//
// Mapzen
// Melbourne VIC location is (45.551136, 11.533929)
// Address of (-37.813611,144.963056) is Stop 3: Bourke Street Mall, Bourke Street, Melbourne, Australia
// Detailed address: &geo.Address{FormattedAddress:"Stop 3: Bourke Street Mall, Bourke Street, Melbourne, Australia", Street:"", HouseNumber:"", Suburb:"", Postcode:"", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AUS", City:""}
//
// TomTom
// Melbourne VIC location is (-37.815340, 144.963230)
// Address of (-37.813611,144.963056) is Doyles Road, Elaine, West Central Victoria, Victoria, 3334
// Detailed address: &geo.Address{FormattedAddress:"Doyles Road, Elaine, West Central Victoria, Victoria, 3334", Street:"Doyles Road", HouseNumber:"", Suburb:"", Postcode:"3334", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Elaine"}
//
// Yandex
// Melbourne VIC location is (41.926823, 2.254232)
// Address of (-37.813611,144.963056) is Victoria, City of Melbourne, Elizabeth Street
// Detailed address: &geo.Address{FormattedAddress:"Victoria, City of Melbourne, Elizabeth Street", Street:"Elizabeth Street",
//  HouseNumber:"", Suburb:"", Postcode:"", State:"Victoria", StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU",
//  City:"City of Melbourne"}
//
// FrenchAPIGouv
//  Champs de Mars Paris location is (2.304770, 48.854395)
//  Address of (48.854395,2.304770) is 9001, Parc du Champs de Mars, 75007, Paris, Paris, Île-de-France, France
//  Detailed address: &geo.Address{FormattedAddress:"9001, Parc du Champs de Mars, 75007, Paris, Paris, Île-de-France, France",
//   Street:"Parc du Champs de Mars", HouseNumber:"9001", Suburb:"", Postcode:"75007", State:" Île-de-France", StateDistrict:"",
//   County:" Paris", Country:"France", CountryCode:"FRA", City:"Paris"}
//
// ChainedAPI[OpenStreetmap -> Google]
// Melbourne VIC location is (-37.814217, 144.963161)
// Address of (-37.813611,144.963056) is Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne, Victoria, 3000, Australia
// Detailed address: &geo.Address{FormattedAddress:"Melbourne's GPO, Postal Lane, Chinatown, Melbourne, City of Melbourne, Greater Melbourne,
// 	Victoria, 3000, Australia", Street:"Postal Lane", HouseNumber:"", Suburb:"Melbourne", Postcode:"3000", State:"Victoria",
// 	StateDistrict:"", County:"", Country:"Australia", CountryCode:"AU", City:"Melbourne"}

func try(geocoder geo.Geocoder) { _ = "STUB: not implemented"; return }

func tryOnlyForAmap(geocoder geo.Geocoder) { _ = "STUB: not implemented"; return }

func tryOnlyFRData(geocoder geo.Geocoder) { _ = "STUB: not implemented"; return }
