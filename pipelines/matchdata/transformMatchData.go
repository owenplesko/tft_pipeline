package matchdata

import (
	"oplesko.com/tft_pipeline/riot"
	"oplesko.com/tft_pipeline/types"
)

func TransformMatchData(in chan *riot.RiotTFTMatchResponse) ([]*types.TFTMatch, *types.AugmentStatsArr) {
	matchStream, matchArr := filterNonRankedGames(transformToTFTMatch(in))
	augmentStats := buildTFTAugmentStats(matchStream)
	return matchArr, augmentStats
}

func transformToTFTMatch(in chan *riot.RiotTFTMatchResponse) chan *types.TFTMatch {
	out := make(chan *types.TFTMatch)
	go func() {
		for raw := range in {
			out <- types.NewTFTMatch(raw)
		}
		close(out)
	}()
	return out
}

func filterNonRankedGames(in chan *types.TFTMatch) (chan *types.TFTMatch, []*types.TFTMatch) {
	out := make(chan *types.TFTMatch)
	arr := []*types.TFTMatch{}

	go func() {
		for match := range in {
			if match.QueueId == 1100 {
				out <- match
				arr = append(arr, match)
			}
		}
		close(out)
	}()

	return out, arr
}

func buildTFTAugmentStats(in chan *types.TFTMatch) *types.AugmentStatsArr {
	augmentStats := types.NewAugmentStatsArr()

	for match := range in {
		for _, comp := range match.Comps {
			for i, augmentId := range comp.Augments {
				augmentStats.InsertAugment(types.AugmentOccurence{
					GameDate:    match.Date,
					GameVersion: match.GameVersion,
					AugmentId:   augmentId,
					Pick:        i + 1,
					Tier:        "placeholder",
					Placement:   comp.Placement,
				})
			}
		}
	}
	return augmentStats
}
