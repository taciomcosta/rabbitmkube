package main

// CalculateNewTotalReplicas determines the new number of replicas of a WorkSet
func CalculateNewTotalReplicas(worker WorkerSet) int {
	result := divideCeiling(worker.MessagesOnQueue, worker.MessagesPerPod)
	result = min(result, worker.MaxReplicas)
	return max(result, worker.MinReplicas)
}

func divideCeiling(a, b int) int {
	return (a / b) + (a % b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
