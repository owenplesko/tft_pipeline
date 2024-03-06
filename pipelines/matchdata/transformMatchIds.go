package matchdata

import "oplesko.com/tft_pipeline/database"

func TransformMatchIds(matchOccurences chan MatchOccurence) []MatchOccurence {
	return filterKnownMatchIds(filterDuplicates(matchOccurences))
}

func filterDuplicates(in chan MatchOccurence) chan MatchOccurence {
	out := make(chan MatchOccurence)

	go func() {
		seenStr := make(map[string]bool)

		for occurence := range in {
			if _, seen := seenStr[occurence.MatchId]; !seen {
				seenStr[occurence.MatchId] = true
				out <- occurence
			}
		}

		close(out)
	}()

	return out
}

func filterKnownMatchIds(in chan MatchOccurence) []MatchOccurence {
	unknownMatchIds := []MatchOccurence{}

	for occurence := range in {
		if !database.MatchIdIsKnown(occurence.MatchId) {
			unknownMatchIds = append(unknownMatchIds, occurence)
		}
	}

	return unknownMatchIds
}
