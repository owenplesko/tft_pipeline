package riot

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type RiotRankEntryResponse struct {
	Puuid    string `json:"puuid"`
	Tier     string `json:"tier"`
	Rank     string `json:"rank"`
	Inactive bool   `json:"inactive"`
}

func GetRankEntries(rankEntryPath string, page int) []RiotRankEntryResponse {
	ctx := context.Background()
	url := fmt.Sprintf("https://na1.api.riotgames.com%v?queue=RANKED_TFT&page=%v", rankEntryPath, page)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Riot-Token", os.Getenv("RIOT_API_KEY"))

	regionalRateLimit.Wait(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return nil
	}
	defer resp.Body.Close()

	// Check if the request was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %v\n", resp.StatusCode)
		return nil
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil
	}

	// Parse the JSON data into a slice of TFTRankEntry
	var entries []RiotRankEntryResponse
	err = json.Unmarshal(body, &entries)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return nil
	}

	return entries
}
