package rankedplayers

import (
	"oplesko.com/tft_pipeline/riot"
	"oplesko.com/tft_pipeline/types"
)

func Transform(in chan *riot.RiotRankEntryResponse) chan *types.TFTRankedPlayer {
	return transformToRankedPlayer(in)
}

func transformToRankedPlayer(in chan *riot.RiotRankEntryResponse) chan *types.TFTRankedPlayer {
	out := make(chan *types.TFTRankedPlayer)

	go func() {
		for riotRes := range in {
			out <- types.NewTFTRankedPlayer(riotRes)
		}

		close(out)
	}()

	return out
}
