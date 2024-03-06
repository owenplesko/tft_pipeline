package matchdata

import "oplesko.com/tft_pipeline/riot"

func ExtractMatchData(matchOccurences []MatchOccurence) chan MatchResponseWithContext {
	return produceTFTMatchData(matchOccurences)
}

type MatchResponseWithContext struct {
	Match    *riot.RiotTFTMatchResponse
	GameTier string
}

func produceTFTMatchData(matchOccurences []MatchOccurence) chan MatchResponseWithContext {
	out := make(chan MatchResponseWithContext)

	go func() {
		for _, occurence := range matchOccurences {
			if matchRes, err := riot.RequestTFTMatchData(occurence.MatchId); err == nil {
				out <- MatchResponseWithContext{
					Match:    matchRes,
					GameTier: occurence.GameTier,
				}
			}
		}
		close(out)
	}()

	return out
}
