package database

import (
	"database/sql"
	"log"
	"time"

	"oplesko.com/tft_pipeline/types"
)

func StoreTFTMatch(tx *sql.Tx, match *types.TFTMatch) {
	_, err := tx.Exec(`
		INSERT INTO tft_match (match_id, date_played)
		VALUES ($1, $2)
	`, match.MatchId, match.Date)
	if err != nil {
		log.Printf("Error storing match to db: %v\n", err)
	}
}

func MatchIdIsKnown(matchId string) bool {
	var exists bool

	err := db.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM tft_match WHERE match_id = $1)
	`, matchId).Scan(&exists)
	if err != nil {
		log.Println("error checking if matchId is known")
		return false
	}

	return exists
}

func PruneMatchIds(daysOld time.Duration) {
	daysAgo := time.Now().Add(-daysOld)

	_, err := db.Exec(`
		DELETE FROM tft_match
		WHERE date_played < $1
	`, daysAgo)
	if err != nil {
		log.Println("error pruning matchIds: ", err)
	}
}
