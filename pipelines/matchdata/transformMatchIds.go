package matchdata

import "oplesko.com/tft_pipeline/database"

func TransformMatchIds(matchIds chan string) []string {
	return filterKnownMatchIds(filterDuplicates(matchIds))
}

func filterDuplicates(in chan string) chan string {
	out := make(chan string)

	go func() {
		seenStr := make(map[string]bool)

		for str := range in {
			if _, seen := seenStr[str]; !seen {
				seenStr[str] = true
				out <- str
			}
		}

		close(out)
	}()

	return out
}

func filterKnownMatchIds(in chan string) []string {
	unknownMatchIds := []string{}

	for matchId := range in {
		if !database.MatchIdIsKnown(matchId) {
			unknownMatchIds = append(unknownMatchIds, matchId)
		}
	}

	return unknownMatchIds
}
