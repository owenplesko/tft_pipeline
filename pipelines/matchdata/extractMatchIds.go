package matchdata

import (
	"time"

	"oplesko.com/tft_pipeline/database"
	"oplesko.com/tft_pipeline/riot"
	"oplesko.com/tft_pipeline/types"
)

type MatchOccurence struct {
	MatchId  string
	GameTier string
}

func ExtractMatchIds() (chan MatchOccurence, []*types.TFTRankedPlayer) {
	players := extractRankedPlayersToUpdate()
	return getTFTMatchIds(players), players
}

func extractRankedPlayersToUpdate() []*types.TFTRankedPlayer {
	return database.QueryTFTRankedPlayersByMatchesLastUpdated(25)
}

func getTFTMatchIds(rankedPlayers []*types.TFTRankedPlayer) chan MatchOccurence {
	matchIdStream := make(chan MatchOccurence)

	go func() {
		for _, player := range rankedPlayers {
			player.MatchesLastRequested = time.Now()

			matchesAfter := time.Now().Add(-3 * 24 * time.Hour).Unix()
			matchIds := riot.RequestTFTMatchHistory(player.Puuid, matchesAfter)
			for _, matchId := range matchIds {
				matchIdStream <- MatchOccurence{
					MatchId:  matchId,
					GameTier: player.Tier,
				}
			}
		}
		close(matchIdStream)
	}()

	return matchIdStream
}
