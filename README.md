# pod-controller
A Dynamic Kubernetes Informer, which informs about the events happening on pods

## Steps to Run
1. Set Environment variable 'KUBECONFIG' - path to your kubeconfig
2. `gvr, _ := schema.ParseResourceArg("pods.v1.")` -> customise the ParseResourceArg Parameter to your desired resource
3. `go run main.go`