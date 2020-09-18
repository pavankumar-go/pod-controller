package pkg

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/afiskon/promtail-client/promtail"
	c "github.com/pod-controller/promtail"

	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// LogAddEvent logs add events happening in cluster
func LogAddEvent(obj interface{}) {
	u := obj.(*unstructured.Unstructured)
	loki := c.GetClient()

	log(obj, loki, u)                                                                                     // pushes to loki
	logrus.WithFields(logrus.Fields{"event": "AddEvent"}).Info("received add event!, log pushed to loki") // to stdout
}

// LogUpdateEvent logs update events happening in cluster
func LogUpdateEvent(oldObj, obj interface{}) {
	if !reflect.DeepEqual(oldObj, obj) {
		u := obj.(*unstructured.Unstructured)
		loki := c.GetClient()

		log(obj, loki, u)                                                                                           // pushes to loki
		logrus.WithFields(logrus.Fields{"event": "UpdateEvent"}).Info("received update event!, log pushed to loki") // to stdout
	} else {
		logrus.Info("no update event!")
	}
}

// LogDeleteEvent logs delete events happening in cluster
func LogDeleteEvent(obj interface{}) {
	u := obj.(*unstructured.Unstructured)
	loki := c.GetClient()

	log(obj, loki, u)                                                                                           // pushes to loki
	logrus.WithFields(logrus.Fields{"event": "DeleteEvent"}).Info("received delete event!, log pushed to loki") // to stdout
}

// log feeds logs to loki
func log(obj interface{}, loki promtail.Client, u *unstructured.Unstructured) {
	subStrs := []string{"Failed", "Fail", "Err", "Invalid", "Unhealthy", "Not"}
	reason := ExtractNestedString(u, "reason")
	for _, s := range subStrs {
		if strings.Contains(reason, s) && ExtractNestedString(u, "type") == "Warning" {
			loki.Errorf(getLogString(u))
		} else if !strings.Contains(reason, s) && ExtractNestedString(u, "type") == "Warning" {
			loki.Warnf(getLogString(u))
		} else {
			loki.Infof(getLogString(u))
		}
	}
}

// getLogString returns formatted log string
func getLogString(u *unstructured.Unstructured) string {
	return fmt.Sprintf("Namespace = %s, Object = %s, Time = %s, Type = %s, Reason= %s, Message=%s, Count = %d \n",
		u.GetNamespace(),
		ExtractNestedString(u, "involvedObject", "name"),
		ExtractNestedString(u, "lastTimestamp"),
		ExtractNestedString(u, "type"),
		ExtractNestedString(u, "reason"),
		ExtractNestedString(u, "message"),
		ExtractNestedInt(u, "count"))
}
