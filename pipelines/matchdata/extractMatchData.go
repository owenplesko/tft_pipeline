package matchdata

import "oplesko.com/tft_pipeline/riot"

func ExtractMatchData(matchIds []string) chan *riot.RiotTFTMatchResponse {
	return produceTFTMatchData(matchIds)
}

func produceTFTMatchData(matchIds []string) chan *riot.RiotTFTMatchResponse {
	out := make(chan *riot.RiotTFTMatchResponse)

	go func() {
		for _, matchId := range matchIds {
			if matchRes, err := riot.RequestTFTMatchData(matchId); err == nil {
				out <- matchRes
			}
		}
		close(out)
	}()

	return out
}
