package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/golang/glog"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func getResourceAge(createtime time.Time) string {
	t := time.Now()
	age := strings.SplitAfter(t.Sub(createtime).String(), "h")

	return age[0]
}

func listDeployments(client *kubernetes.Clientset, filter metav1.ListOptions) {
	deploymentsClient := client.AppsV1().Deployments(apiv1.NamespaceDefault)

	list, err := deploymentsClient.List(filter)

	if err != nil {
		glog.Error("Unable to list deployments")
		panic(err)
	}

	fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
	for _, d := range list.Items {
		fmt.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
}

func deleteDeployments(client *kubernetes.Clientset, filter metav1.ListOptions, label string) {
	deploymentsClient := client.AppsV1().Deployments(apiv1.NamespaceDefault)

	list, err := deploymentsClient.List(filter)

	if err != nil {
		glog.Error("Unable to list deployments")
		panic(err)
	}

	if len(list.Items) != 0 {
		glog.Info("Deployment list not empty")
		fmt.Printf("Listing deployments in namespace %q:\n", apiv1.NamespaceDefault)
		for _, d := range list.Items {
			podlife := d.Labels[label]
			createtime := d.CreationTimestamp.Time

			if podage := getResourceAge(createtime); podage > podlife {
				fmt.Printf("Deleting Deployment - %s ..\n", d.Name)
				deletePolicy := metav1.DeletePropagationForeground
				if err := deploymentsClient.Delete(d.Name, &metav1.DeleteOptions{PropagationPolicy: &deletePolicy}); err != nil {
					panic(err)
				}

				fmt.Printf("Deleted")

			} else {
				fmt.Printf("Pod age is less than ttl value : &s ..\n", d.Name)
			}
		}
	} else {
		fmt.Printf("No deployment with label")
	}
}

func main() {
	var kubeconfig *string
	var configpath string

	label := "life"
	home := homedir.HomeDir()

	if home != "" {
		configpath = filepath.Join(home, ".kube", "config")
	}

	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		kubeconfig = flag.String("kubeconfig", "", "use config from cluster")
	} else {
		kubeconfig = flag.String("kubeconfig", configpath, "abs path to kubeconfig")
	}

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	optionList := metav1.ListOptions{
		LabelSelector: label,
	}

	listDeployments(clientset, optionList)
	deleteDeployments(clientset, optionList, label)
}
