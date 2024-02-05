package types

import (
	"time"

	"oplesko.com/tft_pipeline/riot"
)

type TFTRankedPlayer struct {
	Puuid                string `json:"puuid"`
	Tier                 string `json:"tier"`
	Rank                 string `json:"rank"`
	Inactive             bool   `json:"inactive"`
	MatchesLastRequested time.Time
}

func NewTFTRankedPlayer(raw *riot.RiotRankEntryResponse) *TFTRankedPlayer {
	return &TFTRankedPlayer{
		Puuid:    raw.Puuid,
		Tier:     raw.Tier,
		Rank:     raw.Rank,
		Inactive: raw.Inactive,
	}
}
