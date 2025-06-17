package main

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//parse the kubeconfig path from flag
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("❌ Cannot determine home directory: %v", err)
	}

	kubeconfig := flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "Path to the kubeconfig file")
	flag.Parse()
	//Build K8s config from kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatal("Failed to load kubeconfig: %v", err)
	}

	//create the kubernetes clinetset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal("Failed to create clientset: %v", err)
	}

	ctx := context.Background()

	// List pods

	fmt.Println("Listing pods in 'default' namespace")
	pods, err := clientset.CoreV1().Pods("default").List(ctx, v1.ListOptions{})
	if err != nil {
		log.Fatal("Failed to list pods: %v", err)
	}
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}

	// List deployments
	fmt.Println("Listing pods in 'kube-system' namespace")
	deployments, err := clientset.AppsV1().Deployments("kube-system").List(ctx, v1.ListOptions{})
	if err != nil {
		log.Fatal("Failed to list deployments: %v", err)
	}
	for _, deploy := range deployments.Items {
		fmt.Printf("Deployment name: %s\n", deploy.Name)
	}
}
