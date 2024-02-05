package rankedplayers

import (
	"oplesko.com/tft_pipeline/database"
	"oplesko.com/tft_pipeline/types"
)

func Load(in chan *types.TFTRankedPlayer) {
	sinkTFTRankedPlayers(in)
}

func sinkTFTRankedPlayers(in chan *types.TFTRankedPlayer) {
	for player := range in {
		database.UpsertTFTRankedPlayer(player)
	}
}
