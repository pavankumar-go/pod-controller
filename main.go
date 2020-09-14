package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"time"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // Needed for gcp auth
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// getConfig returns config from path set in KUBECONFIG env variable
func getConfig() (*rest.Config, error) {
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: os.Getenv("KUBECONFIG")},
		&clientcmd.ConfigOverrides{
			ClusterInfo: api.Cluster{
				Server: "",
			},
			CurrentContext: "",
		}).ClientConfig()
}

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatalf("Failed to get cluster config : %v", err)
		os.Exit(2)
	}

	// using dynamic client
	dynamicClient := dynamic.NewForConfigOrDie(config)

	logrus.Info("Starting Informer")

	informer := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicClient, 0, v1.NamespaceDefault, nil)

	gvr, _ := schema.ParseResourceArg("pods.v1.") // `(resource.group.com)` -> `group=com, version=group, resource=resource` and `group=group.com, resource=resource`

	i := informer.ForResource(*gvr)

	// event hanlders for ADD, DELETE, UPDATE
	handlers := cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			u := obj.(*unstructured.Unstructured)

			logrus.WithFields(logrus.Fields{
				"name":      u.GetName(),
				"namespace": u.GetNamespace(),
				"labels":    u.GetLabels(),
			}).Info("received add event!")
			logrus.Info("Updating Pod..")
			updatePod(dynamicClient, u.GetName())
		},

		UpdateFunc: func(oldObj, obj interface{}) {
			if !reflect.DeepEqual(oldObj, obj) {
				u := obj.(*unstructured.Unstructured)
				logrus.WithFields(logrus.Fields{
					"name":      u.GetName(),
					"namespace": u.GetNamespace(),
					"labels":    u.GetLabels(),
				}).Info("received update event!")
			} else {
				log.Println("not updated")
			}
		},

		DeleteFunc: func(obj interface{}) {
			u := obj.(*unstructured.Unstructured)

			logrus.WithFields(logrus.Fields{
				"name":      u.GetName(),
				"namespace": u.GetNamespace(),
				"labels":    u.GetLabels(),
			}).Info("received delete event!")
		},
	}

	i.Informer().AddEventHandler(handlers)

	stopCh := make(chan struct{})
	defer close(stopCh)
	logrus.Info("Informer Started to Listen for events on pods")
	go i.Informer().Run(stopCh)

	sigCh := make(chan os.Signal, 0)
	signal.Notify(sigCh, os.Kill, os.Interrupt)
	<-sigCh
}

// deletePod deletes pod
func deletePod(client dynamic.Interface, podname string) {
	podRes := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
	deletePolicy := v1.DeletePropagationForeground
	deleteOptions := v1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}
	if err := client.Resource(podRes).Namespace(v1.NamespaceDefault).Delete(podname, &deleteOptions); err != nil {
		panic(err)
	}

	client.Resource(podRes).Namespace(v1.NamespaceDefault).Delete(podname, &deleteOptions)
	fmt.Println("Deleted Pod " + podname)
}

// updatePod patches pod labels
func updatePod(client dynamic.Interface, podname string) {
	podRes := schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}

	payload := []patchStringValue{{
		Op:    "replace",
		Path:  "/metadata/labels/run",
		Value: time.Now().Format("2006-01-02_15.04.05"),
	}}

	payloadBytes, _ := json.Marshal(payload)

	result, err := client.Resource(podRes).Namespace(v1.NamespaceDefault).Patch(podname, types.JSONPatchType, payloadBytes, v1.PatchOptions{})
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.WithFields(logrus.Fields{
		"name":   podname,
		"labels": result.GetLabels(),
	}).Info("Label updated")
}

// patchStringValue data for patching
type patchStringValue struct {
	Op    string `json:"op"`
	Path  string `json:"path"`
	Value string `json:"value"`
}
