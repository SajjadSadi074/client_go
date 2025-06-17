package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	// this is for a program that is running inside a pod
	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
	}

	//creates a kubernetes clientset that allows different resorces to be used
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
	deployments, err := clientset.AppsV1().Deployments("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	}
	for _, deployment := range deployments.Items {
		fmt.Println(deployment.Name)
	}
}
