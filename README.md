# pod-controller
A Dynamic Kubernetes Informer, which informs about the events happening on pods

## Steps to Run
1. Set Environment variable 'KUBECONFIG' - path to your kubeconfig
2. `gvr, _ := schema.ParseResourceArg("pods.v1.")` -> change the ParseResourceArg Parameter to your desired GVR `(resource.group.com)` -> `group=com, version=group, resource=resource` and `group=group.com, resource=resource`
3. `make build`