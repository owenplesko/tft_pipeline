package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"oplesko.com/tft_pipeline/cron"
	"oplesko.com/tft_pipeline/database"
	"oplesko.com/tft_pipeline/pipelines/matchdata"
	"oplesko.com/tft_pipeline/riot"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load .env file")
	}

	database.InitConnection(os.Getenv("DB_CONNECTION"))
	riot.InitRateLimit(95, 2*time.Minute)

	go matchdata.BeginMatchDataPipeline()

	cron.RunCronTasks()
}
