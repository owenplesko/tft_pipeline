package types

import (
	"time"
)

type AugmentStat struct {
	GameDate              time.Time
	GameVersion           string
	AugmentId             string
	AggregateAugmentStats []AggregateAugmentStat
}

type AggregateAugmentStat struct {
	TotalPlacement int
	Frequency      int
}

type AugmentOccurence struct {
	GameDate    time.Time
	GameVersion string
	AugmentId   string
	Pick        int
	Placement   int
}

type AugmentStatsArr []*AugmentStat

func NewAugmentStatsArr() *AugmentStatsArr {
	return new(AugmentStatsArr)
}

func (this *AugmentStatsArr) InsertAugment(augmentOccurence AugmentOccurence) {
	if this == nil {
		return
	}

	for _, stat := range *this {
		if stat.GameDate == augmentOccurence.GameDate &&
			stat.GameVersion == augmentOccurence.GameVersion &&
			stat.AugmentId == augmentOccurence.AugmentId {

			stat.AggregateAugmentStats[0].Frequency += 1
			stat.AggregateAugmentStats[0].TotalPlacement += augmentOccurence.Placement
			stat.AggregateAugmentStats[augmentOccurence.Pick].Frequency += 1
			stat.AggregateAugmentStats[augmentOccurence.Pick].TotalPlacement += augmentOccurence.Placement
			return
		}
	}

	stat := AugmentStat{
		GameDate:              augmentOccurence.GameDate,
		GameVersion:           augmentOccurence.GameVersion,
		AugmentId:             augmentOccurence.AugmentId,
		AggregateAugmentStats: make([]AggregateAugmentStat, 4),
	}
	stat.AggregateAugmentStats[0].Frequency += 1
	stat.AggregateAugmentStats[0].TotalPlacement += augmentOccurence.Placement
	stat.AggregateAugmentStats[augmentOccurence.Pick].Frequency += 1
	stat.AggregateAugmentStats[augmentOccurence.Pick].TotalPlacement += augmentOccurence.Placement

	*this = append(*this, &stat)
}
