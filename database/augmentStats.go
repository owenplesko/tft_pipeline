package database

import (
	"database/sql"

	"oplesko.com/tft_pipeline/types"
)

func InsertOrUpdateAugmentStats(tx *sql.Tx, augmentStats *types.AugmentStatsArr) error {
	for _, stat := range *augmentStats {
		// Use INSERT ... ON CONFLICT to handle conflicts and update if necessary
		_, err := tx.Exec(`
			INSERT INTO tft_augment_stats (match_date, game_version, augment_id, pick, accumulated_placement, frequency)
			VALUES ($1, $2, $3, $4, $5, $6)
			ON CONFLICT (match_date, game_version, augment_id, pick)
			DO UPDATE SET
				accumulated_placement = tft_augment_stats.accumulated_placement + $5,
				frequency = tft_augment_stats.frequency + $6
		`,
			stat.GameDate, stat.GameVersion, stat.AugmentId, stat.Pick, stat.AccumulatedPlacement, stat.Frequency)
		if err != nil {
			return err
		}
	}

	return nil
}
