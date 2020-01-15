package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	log.Println("test")
	db, err := sql.Open("mysql", "root:secretpassword@tcp(localhost:30099)/my-database")
	log.Println(err)
	log.Println(db)

	// uses the current context in kubeconfig
	// path-to-kubeconfig -- for example, /root/.kube/config
	config, _ := clientcmd.BuildConfigFromFlags("", "/home/vagrant/.kube/config")

	// ================================================
	// When we move this to inside K8s cluster :)
	var _, _ = rest.InClusterConfig()

	// creates the clientset
	clientset, _ := kubernetes.NewForConfig(config)
	watch, _ := clientset.EventsV1beta1().Events("onap").Watch(v1.ListOptions{})

	// access the API to list pods
	sts, _ := clientset.AppsV1().StatefulSets("onap").Get("mariadb-galera-mariadb-galera", v1.GetOptions{})
	log.Println(sts.Status.ReadyReplicas)

	for x := range watch.ResultChan() {
		log.Println(x.Type + "   |")
		// log.Println(x)

		unstructuredObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(x.Object)
		if err != nil {
			log.Println(err)
		}
		log.Println(unstructuredObj["reason"])
		log.Println(unstructuredObj["type"])
		x, _ := json.Marshal(unstructuredObj)
		log.Println(string(x))
		log.Println()
		log.Println()
		log.Println()
	}
}
