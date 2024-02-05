package matchdata

import (
	"time"

	"oplesko.com/tft_pipeline/database"
	"oplesko.com/tft_pipeline/riot"
	"oplesko.com/tft_pipeline/types"
)

func ExtractMatchIds() (chan string, []*types.TFTRankedPlayer) {
	players := extractRankedPlayersToUpdate()
	return getTFTMatchIds(players), players
}

func extractRankedPlayersToUpdate() []*types.TFTRankedPlayer {
	return database.QueryTFTRankedPlayersByMatchesLastUpdated(10)
}

func getTFTMatchIds(rankedPlayers []*types.TFTRankedPlayer) chan string {
	matchIdStream := make(chan string)

	go func() {
		for _, player := range rankedPlayers {
			player.MatchesLastRequested = time.Now()

			matchesAfter := time.Now().Add(-3 * 24 * time.Hour).Unix()
			matchIds := riot.RequestTFTMatchHistory(player.Puuid, matchesAfter)
			for _, matchId := range matchIds {
				matchIdStream <- matchId
			}
		}
		close(matchIdStream)
	}()

	return matchIdStream
}
