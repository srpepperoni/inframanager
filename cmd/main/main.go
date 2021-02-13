package main

import (
	"context"
	"fmt"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/getPods", GetPodsFromNamespace)
	mux.HandleFunc("/getDeplois", GetDeployments)
	http.ListenAndServe(":9090", mux)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hola mundo")
}

func GetPodsFromNamespace(w http.ResponseWriter, r *http.Request) {
	config, err := rest.InClusterConfig()

	fmt.Fprint(w, "Config ServerName\n", config.ServerName, "\n", config.ContentConfig)

	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})

	fmt.Fprint(w, "There are d pods in the cluster\n", len(pods.Items))
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

	dep, err := clientset.AppsV1().Deployments("").Get(context.TODO(), "prueba-go", metav1.GetOptions{})

	fmt.Fprintln(w, dep.Name)

}
