package rankedplayers

import "time"

func BeginRankedPlayerPipeline() {
	for range time.Tick(24 * time.Hour) {
		Load(Transform(Extract()))
	}
}
