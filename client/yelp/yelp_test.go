package yelp

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewYelpClient(t *testing.T) {
	want := YelpClient{http.DefaultClient, "https://api.yelp.com/v3", "test key"}
	got := NewYelpClient("test key")
	if got != want {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}

func TestSearchBusiness(t *testing.T) {
	yc := NewYelpClient("test key")

	yc.client.Transport = RoundTripper(func (req *http.Request) *http.Response {
		expectedUrl := "https://api.yelp.com/v3/businesses/search?queryUrl"
		actualUrl := req.URL.String()
		if expectedUrl != actualUrl {
			t.Errorf("expected Url: %s, got: %s", expectedUrl, actualUrl)
		}
		expectedAuthHeader := "Bearer test key"
		actualAuthHeader := req.Header.Get("Authorization")
		if actualAuthHeader != expectedAuthHeader {
			t.Errorf("expected Authorization: %s, got: %s", expectedAuthHeader, actualAuthHeader)
		}
		return &http.Response{
			Header: make(http.Header),
			StatusCode: 200,
			Body: io.NopCloser(bytes.NewBufferString(`{"businesses":[{"id":"id","alias":"alias","name":"name","phone":"+15558337200","distance":81.97899441846013,"categories":[{"alias":"alias","title":"title"}],"coordinates":{"latitude":45.02417,"longitude":-70.22048},"location":{"address1":"addr1","address2":"","address3":"","city":"city","country":"US","state":"SA","zip_code":"12345"}}],"total":766}`)),
		}
	})
	
	want := SearchBusinessesResult{
		Total: 766,
		Businesses: []BusinessResult{
			{
				Id: "id",
				Alias: "alias",
				Name: "name",
				Phone: "+15558337200",
				Distance: 81.97899441846013,
				Categories: []BusinessResultCategory{
					{
						Alias: "alias",
						Title: "title",
					},
				},
				Coordinates: BusinessResultCoordinates{
					Latitude: 45.02417,
					Longitude: -70.22048,
				},
				Location: BusinessResultLocation{
					Address1: "addr1",
					Address2: "",
					Address3: "",
					City: "city",
					Country: "US",
					State: "SA",
					ZipCode: "12345",
				},
			},
		},
	}
	got, _ := yc.SearchBusinesses("?queryUrl")
	if !cmp.Equal(got, want) {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}


type RoundTripper func(req *http.Request) *http.Response

func (f RoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}