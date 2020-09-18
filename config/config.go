package config

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// GetClientSet returns dynamic client set for the config, a path set in KUBECONFIG as env variable
func GetClientSet() dynamic.Interface {
	config := loadConfig()
	logrus.Info("Creating New Dynamic Interface for Rest Config...")
	return dynamic.NewForConfigOrDie(config)
}

func loadConfig() *rest.Config {
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		config, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalf("Failed to get config : %v", err)
		}
		return config
	}

	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig},
		&clientcmd.ConfigOverrides{
			ClusterInfo: api.Cluster{
				Server: "",
			},
			CurrentContext: "", // leave blank to use default context in kubeconfig
		}).ClientConfig()
	if err != nil {
		log.Fatalf("Failed to get config : %v", err)
	}

	return config
}
