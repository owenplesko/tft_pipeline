package types

import (
	"time"
)

type AugmentStat struct {
	GameDate             time.Time
	GameVersion          string
	AugmentId            string
	Pick                 int
	AccumulatedPlacement int
	Frequency            int
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
			stat.AugmentId == augmentOccurence.AugmentId &&
			stat.Pick == augmentOccurence.Pick {
			stat.AccumulatedPlacement += augmentOccurence.Placement
			stat.Frequency += 1
			return
		}
	}

	*this = append(*this, &AugmentStat{
		GameDate:             augmentOccurence.GameDate,
		GameVersion:          augmentOccurence.GameVersion,
		AugmentId:            augmentOccurence.AugmentId,
		Pick:                 augmentOccurence.Pick,
		AccumulatedPlacement: augmentOccurence.Placement,
		Frequency:            1,
	})
}
