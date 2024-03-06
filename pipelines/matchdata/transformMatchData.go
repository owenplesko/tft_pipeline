package matchdata

import (
	"oplesko.com/tft_pipeline/types"
)

func TransformMatchData(in chan MatchResponseWithContext) ([]*types.TFTMatch, *types.AugmentStatsArr) {
	matches := transformToTFTMatch(in)
	rankedMatches := filterNonRankedGames(matches)
	augmentStats := buildTFTAugmentStats(rankedMatches)
	return matches, augmentStats
}

func transformToTFTMatch(in chan MatchResponseWithContext) []*types.TFTMatch {
	arr := []*types.TFTMatch{}
	for matchRes := range in {
		arr = append(arr, types.NewTFTMatch(matchRes.Match, matchRes.GameTier))
	}
	return arr
}

func filterNonRankedGames(in []*types.TFTMatch) []*types.TFTMatch {
	arr := []*types.TFTMatch{}

	for _, match := range in {
		if match.QueueId == 1100 {
			arr = append(arr, match)
		}
	}

	return arr
}

func buildTFTAugmentStats(arr []*types.TFTMatch) *types.AugmentStatsArr {
	augmentStats := types.NewAugmentStatsArr()

	for _, match := range arr {
		for _, comp := range match.Comps {
			for i, augmentId := range comp.Augments {
				augmentStats.InsertAugment(types.AugmentOccurence{
					GameTier:    match.GameTier,
					GameVersion: match.GameVersion,
					AugmentId:   augmentId,
					Pick:        i + 1,
					Placement:   comp.Placement,
				})
			}
		}
	}
	return augmentStats
}
