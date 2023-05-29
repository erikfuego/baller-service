package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	bytes := makeGetRequest("https://stats.nba.com/stats/commonallplayers?IsOnlyCurrentSeason=0&LeagueID=00&Season=2023")
	
	fmt.Printf("%d\n", len(bytes))
}

func makeGetRequest(url string) []byte {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Host", "stats.nba.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	req.Header.Add("x-nba-stats-origin", "stats")
	req.Header.Add("x-nba-stats-token", "true")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://stats.nba.com/")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")
	fmt.Printf("HTTP GET %s\n", url)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status: %s\n", res.Status)
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return bytes
}
