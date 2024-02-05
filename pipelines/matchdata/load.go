package matchdata

import (
	"database/sql"
	"log"

	"oplesko.com/tft_pipeline/database"
	"oplesko.com/tft_pipeline/types"
)

func Load(rankedPlayers []*types.TFTRankedPlayer, matches []*types.TFTMatch, augmentStats *types.AugmentStatsArr) {
	tx, _ := database.NewTransaction()
	sinkMatchesLastRequested(tx, rankedPlayers)
	sinkTFTMatches(tx, matches)
	sinkAugmentStats(tx, augmentStats)
	err := tx.Commit()
	if err != nil {
		log.Printf("failed to commit tx: %v\n", err)
	}
}

func sinkMatchesLastRequested(tx *sql.Tx, rankedPlayers []*types.TFTRankedPlayer) {
	for _, player := range rankedPlayers {
		database.UpdateMatchesLastRequested(tx, player)
	}
}

func sinkTFTMatches(tx *sql.Tx, matches []*types.TFTMatch) {
	for _, match := range matches {
		database.StoreTFTMatch(tx, match)
	}
}

func sinkAugmentStats(tx *sql.Tx, augmentStats *types.AugmentStatsArr) {
	database.InsertOrUpdateAugmentStats(tx, augmentStats)
}
