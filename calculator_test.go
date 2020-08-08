package main

import "testing"

type TestCase struct {
	want        int
	input       WorkerSet
	description string
}

var tests []TestCase = []TestCase{
	{
		description: "MessagesOnQueue=5, MessagesPerPod=1 => 5 replicas",
		want:        5,
		input: WorkerSet{
			MessagesOnQueue: 5,
			MinReplicas:     3,
			MaxReplicas:     8,
			MessagesPerPod:  1,
		},
	},
	{
		description: "MessagesOnQueue=5, MessagesPerPod=5 => 1 replicas",
		want:        1,
		input: WorkerSet{
			MessagesOnQueue: 5,
			MinReplicas:     1,
			MaxReplicas:     8,
			MessagesPerPod:  5,
		},
	},
	{
		description: "MessagesOnQueue=5, MessagesPerPod=2 => 3 replicas",
		want:        3,
		input: WorkerSet{
			MessagesOnQueue: 5,
			MinReplicas:     3,
			MaxReplicas:     8,
			MessagesPerPod:  2,
		},
	},
	{
		description: "MessagesOnQueue=0, MinReplicas=1 => 1 replica",
		want:        1,
		input: WorkerSet{
			MessagesOnQueue: 0,
			MinReplicas:     1,
			MaxReplicas:     10,
			MessagesPerPod:  1,
		},
	},
	{
		description: "MessagesOnQueue=20, MaxReplicas=10, MessagesPerPod=1 => 10 replicas",
		want:        10,
		input: WorkerSet{
			MessagesOnQueue: 20,
			MinReplicas:     1,
			MaxReplicas:     10,
			MessagesPerPod:  1,
		},
	},
	{
		description: "MessagesOnQueue=5, MessagesPerPod=2, MaxReplicas=2 => 2 replicas",
		want:        2,
		input: WorkerSet{
			MessagesOnQueue: 5,
			MinReplicas:     1,
			MaxReplicas:     2,
			MessagesPerPod:  2,
		},
	},
	{
		description: "MessagesPerPod=4, MessagesOnQueue=20 => 5 replicas",
		want:        5,
		input: WorkerSet{
			MessagesOnQueue: 20,
			MinReplicas:     1,
			MaxReplicas:     10,
			MessagesPerPod:  4,
		},
	},
	{
		description: "MinReplicas=3, MessagesOnQueue=2 => 3 replicas",
		want:        3,
		input: WorkerSet{
			MessagesOnQueue: 2,
			MinReplicas:     3,
			MaxReplicas:     10,
			MessagesPerPod:  4,
		},
	},
	{
		description: "MinReplicas=2, MessagesOnQueue=4, MessagesPerPod=4 => 2 replicas",
		want:        2,
		input: WorkerSet{
			MessagesOnQueue: 4,
			MinReplicas:     2,
			MaxReplicas:     10,
			MessagesPerPod:  4,
		},
	},
}

func TestCalculateNewTotalReplicas(t *testing.T) {
	for _, tt := range tests {
		got := CalculateNewTotalReplicas(tt.input)
		if tt.want != got {
			t.Error(tt.description)
		}
	}
}
