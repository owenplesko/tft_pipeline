package types

import (
	"regexp"
	"time"

	"oplesko.com/tft_pipeline/riot"
)

type TFTMatch struct {
	MatchId     string
	Date        time.Time
	QueueId     int
	GameVersion string
	Comps       []riot.RiotTFTCompResponse
}

func NewTFTMatch(raw *riot.RiotTFTMatchResponse) *TFTMatch {
	match := &TFTMatch{
		MatchId:     raw.MetaData.MatchId,
		Date:        time.UnixMilli(raw.Info.Date).Round(24 * time.Hour),
		QueueId:     raw.Info.QueueId,
		GameVersion: extractGameVersion(raw.Info.Version),
		Comps:       raw.Info.Comps,
	}
	return match
}

var versionRegex = regexp.MustCompile(`<Releases\/(\d+\.\d+)>`)

func extractGameVersion(versionString string) string {
	matches := versionRegex.FindStringSubmatch(versionString)
	gameVersion := matches[1]
	return gameVersion
}
