package main

import (
	"context"
	"fmt"
	"inframanager/internal/datalayer"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/getPods", GetPodsFromNamespace)
	mux.HandleFunc("/getDeploy", GetDeployments)
	http.ListenAndServe(":9090", mux)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func GetPodsFromNamespace(w http.ResponseWriter, r *http.Request) {
	config, err := rest.InClusterConfig()

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintln(w, "There are ", len(pods.Items), " pods in the cluster")

	for index, element := range pods.Items {
		fmt.Fprintln(w, "POD ", index, " : ", element.Name)
	}

}

func GetDeployments(w http.ResponseWriter, r *http.Request) {
	config, err := rest.InClusterConfig()

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	dep, err := clientset.AppsV1().Deployments("default").List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintln(w, "There are ", len(dep.Items), " deployments in the cluster")

	for index, element := range dep.Items {
		fmt.Fprintln(w, "Deployment ", index, ": ", element.Name)
	}

}

func GetPlayersFromMongo() {
	var client = datalayer.InitDataLayer()
	defer client.Disconnect(context.Background())
}
