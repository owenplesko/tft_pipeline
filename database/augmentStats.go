package database

import (
	"database/sql"
	"log"

	"oplesko.com/tft_pipeline/types"
)

func InsertOrUpdateAugmentStats(tx *sql.Tx, augmentStats *types.AugmentStatsArr) error {
	for _, stat := range *augmentStats {
		// Use INSERT ... ON CONFLICT to handle conflicts and update if necessary
		_, err := tx.Exec(`
			INSERT INTO tft_augment_stats (
				match_date,
				game_version,
				augment_id,
				total_placement,
				frequency,
				pick_1_total_placement,
				pick_1_frequency,
				pick_2_total_placement,
				pick_2_frequency,
				pick_3_total_placement,
				pick_3_frequency
			)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
			ON CONFLICT (match_date, game_version, augment_id)
			DO UPDATE SET
				total_placement = tft_augment_stats.total_placement + $4,
				frequency = tft_augment_stats.frequency + $5,
				pick_1_total_placement = tft_augment_stats.pick_1_total_placement + $6,
				pick_1_frequency = tft_augment_stats.pick_1_frequency + $7,
				pick_2_total_placement = tft_augment_stats.pick_2_total_placement + $8,
				pick_2_frequency = tft_augment_stats.pick_2_frequency + $9,
				pick_3_total_placement = tft_augment_stats.pick_3_total_placement + $10,
				pick_3_frequency = tft_augment_stats.pick_3_frequency + $11
		`,
			stat.GameDate,
			stat.GameVersion,
			stat.AugmentId,
			stat.AggregateAugmentStats[0].TotalPlacement,
			stat.AggregateAugmentStats[0].Frequency,
			stat.AggregateAugmentStats[1].TotalPlacement,
			stat.AggregateAugmentStats[1].Frequency,
			stat.AggregateAugmentStats[2].TotalPlacement,
			stat.AggregateAugmentStats[2].Frequency,
			stat.AggregateAugmentStats[3].TotalPlacement,
			stat.AggregateAugmentStats[3].Frequency,
		)
		if err != nil {
			log.Println("error inserting augment stats: ", err)
			return err
		}
	}

	return nil
}
