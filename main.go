package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeConfig *string
	//path of the config file,if path not specified will go search at the default path
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "path of kubeconfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "path to config file")
	}
	flag.Parse()

	//buliding the configuration from the kubernets file
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}

	//creating a client for the use configuration file
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	//getting the pods detatils from the cluster
	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//printing list of items in the pods
	fmt.Print("pods Items :\n", pods.Items)
	//getting the count of the pods from the cluster
	fmt.Printf("there are %d pods in the cluster\n", len(pods.Items))

	//getting the name and namespace of the pods from the cluster
	for _, pod := range pods.Items {
		fmt.Printf("NameSpace: %s, PodName: %s\n", pod.Namespace, pod.Name)
	}
}
