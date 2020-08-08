package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	v1 "k8s.io/api/apps/v1"
)

// WorkerSet wraps v1.Deployment
type WorkerSet struct {
	MessagesOnQueue int
	MinReplicas     int
	MaxReplicas     int
	MessagesPerPod  int
	Deploy          *v1.Deployment
	Queue           string
}

// NewWorkerSetList creates a []WorkerSet based on a []v1.Deployment
func NewWorkerSetList(deployList *v1.DeploymentList) []WorkerSet {
	workerSets := make([]WorkerSet, 0)
	for _, deploy := range deployList.Items {
		if isAnnotated(deploy) {
			workerSets = append(workerSets, NewWorkerSet(deploy.DeepCopy())) // DeepCopy due to mutable state in k8s lib
		}
	}
	return workerSets
}

func isAnnotated(deploy v1.Deployment) bool {
	_, ok := deploy.Annotations["rabbitmkube/min-replicas"]
	return ok
}

// NewWorkerSet creates a WorkerSet based on a k8s Deployment
func NewWorkerSet(deploy *v1.Deployment) WorkerSet {
	var workerSet WorkerSet
	workerSet.MinReplicas, _ = strconv.Atoi(deploy.Annotations["rabbitmkube/min-replicas"])
	workerSet.MaxReplicas, _ = strconv.Atoi(deploy.Annotations["rabbitmkube/max-replicas"])
	workerSet.MessagesPerPod, _ = strconv.Atoi(deploy.Annotations["rabbitmkube/messages-per-pod"])
	workerSet.Queue, _ = deploy.Annotations["rabbitmkube/queue-name"]
	workerSet.Deploy = deploy
	return workerSet
}

// Validate validates if config provided to WorkerSet is ok
func (w WorkerSet) Validate() error {
	if w.MinReplicas < 0 {
		return errors.New("MinReplicas must be >= 0")
	}
	if w.MinReplicas > w.MaxReplicas {
		return errors.New("MinReplicas must be greater or equal to MaxReplicas")
	}
	if w.MessagesPerPod < 1 {
		return errors.New("MessagesPerPod must be >= 1")
	}
	if strings.Trim(w.Queue, " ") == "" {
		return errors.New("Queue cannot be empty")
	}
	return nil
}

// ToString formats WorkerSet to be printed
func (w WorkerSet) ToString() string {
	return fmt.Sprintf(`[%s] has %d replicas and %d messages on queue %s`, w.Deploy.Name, *w.Deploy.Spec.Replicas, w.MessagesOnQueue, w.Queue)
}
