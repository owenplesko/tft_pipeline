package rankedplayers

func BeginRankedPlayerPipeline() {
	for {
		Load(Transform(Extract()))
	}
}
