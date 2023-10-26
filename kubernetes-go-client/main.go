package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	fmt.Println("Starting the program")
	var (
		k8sClient *kubernetes.Clientset
	)

	ctx := context.Background()

	k8sClient, err := getK8sClient()
	if err != nil {
		fmt.Printf("Error is: %s", err)
		os.Exit(1)
	}
	deploymentLables, err := createDeploy(ctx, k8sClient, "default")
	if err != nil {
		fmt.Printf("deployment failed with err : %s", err)
	}
	fmt.Println(deploymentLables)
}

func getK8sClient() (*kubernetes.Clientset, error) {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return clientset, nil
}

func createDeploy(ctx context.Context, k8sClient *kubernetes.Clientset, namespace string) (map[string]string, error) {

	// Function will return labels so map[string]string like name: "Hello World"
	var deployment *v1.Deployment

	deploymentFile, err := os.ReadFile("app.yaml")
	if err != nil {
		return nil, fmt.Errorf("read file error : %s", err)
	}

	obj, groupVersionKind, err := scheme.Codecs.UniversalDeserializer().Decode(deploymentFile, nil, nil) // converts to kubernetes object
	if err != nil {
		return nil, fmt.Errorf("decode failed: %s", err)
	}
	switch val := obj.(type) {
	case *v1.Deployment:
		deployment = val
	default:
		return nil, fmt.Errorf("unrecognixed type : %s", groupVersionKind)
	}
	deploymentResponse, err := k8sClient.AppsV1().Deployments(namespace).Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("deployment error : %s", err)
	}

	return deploymentResponse.Labels, nil
}
