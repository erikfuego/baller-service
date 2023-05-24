package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {
	req, _ := http.NewRequest(http.MethodGet, "https://stats.nba.com/stats/commonplayerinfo?LeagueID=&PlayerID=201939", nil)
	req.Header.Add("Host", "stats.nba.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:72.0) Gecko/20100101 Firefox/72.0")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Accept-Language", "en-US,en;q=0.5")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("x-nba-stats-origin", "stats")
	req.Header.Add("x-nba-stats-token", "true")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://stats.nba.com/")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Cache-Control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("status: %s\n", res.Status)
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	
	fmt.Printf("%s\n", bytes)

}
