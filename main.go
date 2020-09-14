package main

import (
	"log"
	"os"
	"os/signal"
	"reflect"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
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

	log.Println("Starting Informer")

	informer := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicClient, 0, v1.NamespaceDefault, nil)

	gvr, _ := schema.ParseResourceArg("pods.v1.") // pods -> resources, v1-> version, . ->no-group

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
	<-stopCh
}
