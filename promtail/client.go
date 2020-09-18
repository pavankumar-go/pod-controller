package promtail

import (
	"os"
	"time"

	"github.com/afiskon/promtail-client/promtail"
	"github.com/sirupsen/logrus"
)

var (
	promtailClient promtail.Client
	err            error
)

func init() {
	newClient()
}

// NewClient returns promtail client
func newClient() {
	sourceName := os.Getenv("SOURCE_NAME")
	jobName := os.Getenv("JOB_NAME")
	labels := "{source=\"" + sourceName + "\",job=\"" + jobName + "\"}"

	conf := promtail.ClientConfig{
		PushURL:            "http://localhost:3100/api/prom/push", // loki server url, use "<svc-name>.<ns>.svc.cluster.local" -> Incluster URL
		Labels:             labels,
		BatchWait:          5 * time.Second,
		BatchEntriesNumber: 10000,
		SendLevel:          promtail.INFO,
		PrintLevel:         promtail.ERROR,
	}

	promtailClient, err = promtail.NewClientJson(conf)
	if err != nil {
		logrus.Fatalf("Failed to initialise promtail client: %v", err)
	}
}

// GetClient returns promtail client
func GetClient() promtail.Client {
	if promtailClient != nil {
		return promtailClient
	}

	return nil
}
