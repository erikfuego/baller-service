package yelp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewYelpClient(apiKey string) YelpClient {
	return YelpClient{
		client: http.DefaultClient,
		baseUrl: "https://api.yelp.com/v3",
		apiKey: apiKey,
	}
}

func (yc YelpClient) SearchBusinesses(query string) (SearchBusinessesResult, error) {
	endpoint := yc.baseUrl + "/businesses/search"

	req, _ := http.NewRequest("GET", endpoint + query, nil)
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + yc.apiKey)
	
	res, err := yc.client.Do(req)
	if err != nil {
		return SearchBusinessesResult{}, fmt.Errorf("SearchBusiness: failed performing search: %w", err)
	}
	defer res.Body.Close()
	searchBusinessesResult := new(SearchBusinessesResult)
	err = json.NewDecoder(res.Body).Decode(searchBusinessesResult)
	if err != nil {
		return SearchBusinessesResult{}, fmt.Errorf("SearchBusiness: failed to deserialize results: %w", err)
	}
	return *searchBusinessesResult, nil
}

type YelpClient struct {
	client *http.Client
	baseUrl string
	apiKey string
}

type BusinessResultCategory struct {
	Alias string `json:"alias"`
	Title string `json:"title"`
}

type BusinessResultCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type BusinessResultLocation struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Address3 string `json:"address3"`
	City     string `json:"city"`
	Country  string `json:"country"`
	State    string `json:"state"`
	ZipCode  string `json:"zip_code"`
}

type BusinessResult struct {
	Id           string                    `json:"id"`
	Alias        string                    `json:"alias"`
	Name         string                    `json:"name"`
	Phone        string                    `json:"phone"`
	Distance     float64                   `json:"distance"`
	Categories   []BusinessResultCategory  `json:"categories"`
	Coordinates  BusinessResultCoordinates `json:"coordinates"`
	Location     BusinessResultLocation    `json:"location"`
}

type SearchBusinessesResult struct {
	Businesses []BusinessResult `json:"businesses"`
	Total      int              `json:"total"`
}
