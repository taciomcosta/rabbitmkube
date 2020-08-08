package main

import (
	"testing"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestNewWorkerSetList(t *testing.T) {
	workerSet := NewWorkerSetList(mockDeploymentList())
	if len(workerSet) != 1 {
		t.Error("should list deployments with annotations only")
	}
}

func mockDeploymentList() *v1.DeploymentList {
	list := &v1.DeploymentList{}
	list.Items = []v1.Deployment{
		*mockDeployMissingAnnotations(), *mockDeployWithAnnotations(),
	}
	return list
}

func TestNewWorkerSet(t *testing.T) {
	workerSet := NewWorkerSet(mockDeployWithAnnotations())
	if workerSet.MinReplicas != 1 {
		t.Errorf("MinReplicas should be parsed")
	}
	if workerSet.MaxReplicas != 15 {
		t.Errorf("MaxReplicas should be parsed")
	}
	if workerSet.MessagesPerPod != 5 {
		t.Errorf("MessagesPerPod should be parsed")
	}
	if workerSet.Queue != "some-queue" {
		t.Errorf("Queue name should be parsed")
	}
	if workerSet.Deploy == nil {
		t.Errorf("Deploy should be set")
	}
}

func mockDeployWithAnnotations() *v1.Deployment {
	return &v1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"rabbitmkube/min-replicas":     "1",
				"rabbitmkube/max-replicas":     "15",
				"rabbitmkube/messages-per-pod": "5",
				"rabbitmkube/queue-name":       "some-queue",
			},
		},
	}
}

func mockDeployMissingAnnotations() *v1.Deployment {
	return &v1.Deployment{}
}

func TestValidate(t *testing.T) {
	var tests = []struct {
		worker    WorkerSet
		wantError bool
	}{
		{
			worker: WorkerSet{
				MaxReplicas:    0,
				MinReplicas:    1,
				MessagesPerPod: 1,
				Queue:          "queue",
			},
			wantError: true,
		},
		{
			worker: WorkerSet{
				MaxReplicas:    0,
				MinReplicas:    -1,
				MessagesPerPod: 1,
				Queue:          "queue",
			},
			wantError: true,
		},
		{
			worker: WorkerSet{
				MaxReplicas:    0,
				MinReplicas:    0,
				MessagesPerPod: 0,
				Queue:          "queue",
			},
			wantError: true,
		},
		{
			worker: WorkerSet{
				MaxReplicas:    0,
				MinReplicas:    0,
				MessagesPerPod: 1,
				Queue:          " ",
			},
			wantError: true,
		},
	}

	for _, tt := range tests {
		err := tt.worker.Validate()
		if err == nil && tt.wantError {
			t.Errorf("%v should be invalid", tt.worker)
		}
	}
}
