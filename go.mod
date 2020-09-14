module github.com/pod-controller

go 1.13

require (
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kubernetes/client-go v11.0.0+incompatible // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a // indirect
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.19.1 // indirect
	k8s.io/apimachinery v0.19.1
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kube-openapi v0.0.0-20200831175022-64514a1d5d59 // indirect
	k8s.io/utils v0.0.0-20200821003339-5e75c0163111 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)
