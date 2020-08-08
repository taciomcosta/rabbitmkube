package main

import (
	"fmt"
	"time"
)

func main() {
	schedule(scale, time.Duration(10)*time.Second)
}

func scale() {
	queues := FetchQueues()
	for _, workerSet := range ListWorkerSets() {
		scaleOne(workerSet, queues)
	}
}

func scaleOne(workerSet WorkerSet, queues Queues) {
	workerSet.MessagesOnQueue = queues[workerSet.Queue]
	totalReplicas := CalculateNewTotalReplicas(workerSet)
	fmt.Printf("%s, scaling to %d\n", workerSet.ToString(), totalReplicas)
	ScaleWorkerSet(&workerSet, int32(totalReplicas))
}

func schedule(f func(), duration time.Duration) {
	ticker := time.NewTicker(duration)
	for {
		select {
		case <-ticker.C:
			f()
		}
	}
}
