package pkg

import (
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// ExtractNestedString returns the string value of a nested field.
func ExtractNestedString(u *unstructured.Unstructured, fields ...string) string {
	str, ok, err := unstructured.NestedString(u.Object, fields...)
	if !ok || err != nil {
		logrus.Fatalf("Couldn't Extract Nested String Fields : %v , %v", fields, err)
	}

	return str
}

// ExtractNestedInt returns the string value of a nested field.
func ExtractNestedInt(u *unstructured.Unstructured, fields ...string) int64 {
	i64, ok, err := unstructured.NestedInt64(u.Object, fields...)
	if !ok || err != nil {
		logrus.Fatalf("Couldn't Extract Nested Int Field : %v", err)
	}

	return i64
}
