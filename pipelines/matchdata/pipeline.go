package matchdata

func BeginMatchDataPipeline() {
	for {
		matchIdChan, rankedPlayers := ExtractMatchIds()
		matches, tftStats := TransformMatchData(ExtractMatchData(TransformMatchIds(matchIdChan)))
		Load(rankedPlayers, matches, tftStats)
	}
}
