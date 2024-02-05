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

type RiotTFTMatchResponse struct {
	MetaData struct {
		DataVersion  string   `json:"data_version"`
		MatchId      string   `json:"match_id"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		Date      int64                 `json:"game_datetime"`
		Length    float64               `json:"game_length"`
		Version   string                `json:"game_version"`
		QueueId   int                   `json:"queue_id"`
		GameType  string                `json:"tft_game_type"`
		SetName   string                `json:"tft_set_core_name"`
		SetNumber int                   `json:"tft_set_number"`
		Comps     []RiotTFTCompResponse `json:"participants"`
	} `json:"info"`
}

type RiotTFTCompResponse struct {
	Augments  []string `json:"augments"`
	Companion struct {
		ContentId string `json:"content_ID"`
		ItemId    int    `json:"item_ID"`
		SkinId    int    `json:"skin_ID"`
		Species   string `json:"species"`
	} `json:"companion"`
	RemainingGold     int     `json:"gold_left"`
	LastRound         int     `json:"last_round"`
	Level             int     `json:"level"`
	Placement         int     `json:"placement"`
	PlayersEliminated int     `json:"players_eliminated"`
	Puuid             string  `json:"puuid" Match:"Summoner"`
	TimeEliminated    float64 `json:"time_eliminated"`
	DamageToPlayers   int     `json:"total_damage_to_players"`
	Traits            []struct {
		Name       string `json:"name"`
		NumUnits   int    `json:"num_units"`
		Style      int    `json:"style"`
		TierActive int    `json:"tier_current"`
		TierMax    int    `json:"tier_total"`
	} `json:"traits"`
	Units []struct {
		Id        string   `json:"character_id"`
		ItemNames []string `json:"itemNames"`
		Rarity    int      `json:"rarity"`
		Tier      int      `json:"tier"`
	} `json:"units"`
}

func RequestTFTMatchData(matchId string) (*RiotTFTMatchResponse, error) {
	ctx := context.Background()
	url := fmt.Sprintf("https://americas.api.riotgames.com/tft/match/v1/matches/%v", matchId)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Riot-Token", os.Getenv("RIOT_API_KEY"))
	continentalRateLimit.Wait(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Error making request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the request was successful (status code 200)
	if resp.StatusCode != http.StatusOK {
		log.Printf("Request failed with status code: %v\n", resp.StatusCode)
		return nil, err
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, err
	}

	// Parse the JSON data into a slice of TFTRankEntry
	matchData := new(RiotTFTMatchResponse)
	err = json.Unmarshal(body, matchData)
	if err != nil || matchData == nil {
		log.Println("Error decoding JSON:", err)
		return nil, err
	}

	return matchData, nil
}

func RequestTFTMatchHistory(puuid string, matchesAfter int64) []string {
	ctx := context.Background()
	url := fmt.Sprintf("https://americas.api.riotgames.com/tft/match/v1/matches/by-puuid/%v/ids?startTime=%v&count=200", puuid, matchesAfter)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-Riot-Token", os.Getenv("RIOT_API_KEY"))
	continentalRateLimit.Wait(ctx)
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
	var matchIds []string
	err = json.Unmarshal(body, &matchIds)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return nil
	}

	return matchIds
}
