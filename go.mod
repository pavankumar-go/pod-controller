module github.com/pod-controller

go 1.13

require (
	github.com/afiskon/promtail-client v0.0.0-20190305142237-506f3f921e9c
	github.com/containerd/containerd v1.4.0 // indirect
	github.com/containerd/typeurl v1.0.1 // indirect
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/evanphx/json-patch v4.9.0+incompatible // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/jonboulle/clockwork v0.2.1 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/pflag v1.0.5 // indirect
	go.etcd.io/etcd v3.3.25+incompatible // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.19.1 // indirect
	k8s.io/apimachinery v0.19.1
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v1.0.0 // indirect
	k8s.io/kube-openapi v0.0.0-20200831175022-64514a1d5d59 // indirect
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	k8s.io/api => k8s.io/api v0.0.0-20190409021203-6e4e0e4f393b
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190404173353-6a84e37a896d
	k8s.io/client-go => k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
)
