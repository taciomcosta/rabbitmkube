package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var clientset *kubernetes.Clientset

// ScaleWorkerSet updates the number of replicas of a WorkerSet
func ScaleWorkerSet(workerSet *WorkerSet, totalReplicas int32) {
	workerSet.Deploy.Spec.Replicas = &totalReplicas
	_, err := clientset.
		AppsV1().
		Deployments(workerSet.Deploy.Namespace).
		Update(context.TODO(), workerSet.Deploy, metav1.UpdateOptions{})
	if err != nil {
		fmt.Println("Error on scaling WorkerSet")
	}
}

// ListWorkerSets return a []WorkerSets (a list of k8s deployments with rabbitmkube annotations)
func ListWorkerSets() []WorkerSet {
	deploys, err := clientSet().AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error on listing WorkerSets")
	}
	return NewWorkerSetList(deploys)
}

func clientSet() *kubernetes.Clientset {
	if clientset == nil {
		clientset = newClientSet(newInClusterConfig)
	}
	return clientset
}

func newClientSet(config configFunc) *kubernetes.Clientset {
	client, err := kubernetes.NewForConfig(config())
	if err != nil {
		fmt.Println("Error on creating client")
		return nil
	}
	return client
}

type configFunc func() *rest.Config

func newOutClusterConfig() *rest.Config {
	kubeconfig := "/home/taciomcosta/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	return config
}

func newInClusterConfig() *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println("Error on creating in-cluster-config")
		return nil
	}
	return config
}
