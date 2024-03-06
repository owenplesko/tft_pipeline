package database

import (
	"database/sql"
	"log"

	"oplesko.com/tft_pipeline/types"
)

func UpsertTFTRankedPlayer(player *types.TFTRankedPlayer) {

	// Use ON CONFLICT to perform the upsert based on the primary key (puuid)
	query := `
		INSERT INTO ranked_player (puuid, tier, rank, inactive)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (puuid) DO UPDATE
		SET tier = excluded.tier,
		    rank = excluded.rank,
		    inactive = excluded.inactive
	`

	_, err := db.Exec(query, player.Puuid, player.Tier, player.Rank, player.Inactive)
	if err != nil {
		log.Printf("Error storing player to db: %v\n", err)
	}
}

func UpdateMatchesLastRequested(tx *sql.Tx, player *types.TFTRankedPlayer) {
	query := `
		UPDATE ranked_player
		SET matches_last_requested = $2
		WHERE puuid = $1
	`

	_, err := tx.Exec(query, player.Puuid, player.MatchesLastRequested)
	if err != nil {
		log.Printf("Error updating matches_last_updated in db: %v\n", err)
	}
}

func QueryTFTRankedPlayersByMatchesLastUpdated(limit int) []*types.TFTRankedPlayer {
	query := `
		SELECT puuid, tier 
		FROM ranked_player 
		ORDER BY matches_last_requested ASC NULLS FIRST
		LIMIT $1
	`

	rows, err := db.Query(query, limit)
	if err != nil {
		log.Printf("Error querying ranked_player by matches last updated: %v\n", err)
	}

	var rankedPlayers []*types.TFTRankedPlayer
	defer rows.Close()
	for rows.Next() {
		player := &types.TFTRankedPlayer{}
		err := rows.Scan(&player.Puuid, &player.Tier)
		if err != nil {
			log.Printf("Error scanning ranked_player: %v\n", err)
		}

		rankedPlayers = append(rankedPlayers, player)
	}
	return rankedPlayers
}
