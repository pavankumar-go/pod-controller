package controller

import (
	"github.com/pod-controller/config"
	"github.com/pod-controller/pkg"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/dynamicinformer"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // Needed for gcp auth
	"k8s.io/client-go/tools/cache"
)

// GetInformer returns Dynamic Shared Index Informer
func GetInformer() cache.SharedIndexInformer {
	return newInformer()
}

// newInformer function constructs a new instance of dynamic shared informer factory
// adds event handler funcs for k8s events resource and returns shared index informer
func newInformer() cache.SharedIndexInformer {
	dynamicClient := config.GetClientSet()

	informer := dynamicinformer.NewFilteredDynamicSharedInformerFactory(dynamicClient, 0, "", nil)

	gvr, _ := schema.ParseResourceArg("events.v1.") // `(resource.group.com)` -> `group=com, version=group, resource=resource`

	i := informer.ForResource(*gvr)

	i.Informer().AddEventHandler(getHandlers())

	logrus.Info("Starting Informer...")
	return i.Informer()
}

// getHandlers returns events handler funcs
func getHandlers() cache.ResourceEventHandlerFuncs {
	// hanlders for ADD, DELETE, UPDATE events
	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pkg.LogAddEvent(obj) // logs add events
		},

		UpdateFunc: func(oldObj, obj interface{}) {
			pkg.LogUpdateEvent(oldObj, obj) // logs update events
		},

		DeleteFunc: func(obj interface{}) {
			pkg.LogDeleteEvent(obj) // logs delete events
		},
	}
}
