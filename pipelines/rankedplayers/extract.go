package rankedplayers

import (
	"log"

	"oplesko.com/tft_pipeline/riot"
)

var RANK_ENTRY_PATHS = []string{
	//"/tft/league/v1/challenger",
	//"/tft/league/v1/grandmaster",
	//"/tft/league/v1/master",
	"/tft/league/v1/entries/DIAMOND/I",
	"/tft/league/v1/entries/DIAMOND/II",
	"/tft/league/v1/entries/DIAMOND/III",
	"/tft/league/v1/entries/DIAMOND/IV",
	"/tft/league/v1/entries/EMERALD/I",
	"/tft/league/v1/entries/EMERALD/II",
	"/tft/league/v1/entries/EMERALD/III",
	"/tft/league/v1/entries/EMERALD/IV",
	"/tft/league/v1/entries/PLATINUM/I",
	"/tft/league/v1/entries/PLATINUM/II",
	"/tft/league/v1/entries/PLATINUM/III",
	"/tft/league/v1/entries/PLATINUM/IV",
	"/tft/league/v1/entries/GOLD/I",
	"/tft/league/v1/entries/GOLD/II",
	"/tft/league/v1/entries/GOLD/III",
	"/tft/league/v1/entries/GOLD/IV",
	"/tft/league/v1/entries/SILVER/I",
	"/tft/league/v1/entries/SILVER/II",
	"/tft/league/v1/entries/SILVER/III",
	"/tft/league/v1/entries/SILVER/IV",
	"/tft/league/v1/entries/BRONZE/I",
	"/tft/league/v1/entries/BRONZE/II",
	"/tft/league/v1/entries/BRONZE/III",
	"/tft/league/v1/entries/BRONZE/IV",
	"/tft/league/v1/entries/IRON/I",
	"/tft/league/v1/entries/IRON/II",
	"/tft/league/v1/entries/IRON/III",
	"/tft/league/v1/entries/IRON/IV",
}

func Extract() chan *riot.RiotRankEntryResponse {
	return produceTFTRankedPlayers()
}

func produceTFTRankedPlayers() chan *riot.RiotRankEntryResponse {
	out := make(chan *riot.RiotRankEntryResponse)

	go func() {
		for _, entryPath := range RANK_ENTRY_PATHS {
			totalEntries := 0
			page := 1
			for ; ; page++ {
				entries := riot.GetRankEntries(entryPath, page)

				for _, rankEntryRes := range entries {
					out <- &rankEntryRes
				}

				totalEntries += len(entries)

				if len(entries) == 0 {
					break
				}
			}

			log.Printf("finished collecting data from %v with %v pages and %v total entries\n", entryPath, page-1, totalEntries)
		}
		close(out)
	}()

	return out
}
