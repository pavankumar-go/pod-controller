# Informer 
A Dynamic Kubernetes Informer, watches for k8s events and pushes the events as logs to loki (a log monitoring system by Grafana) using promtail client

## Steps to Run
1. Set Environment variable 'KUBECONFIG' - path to your kubeconfig
2. Set Environment variable 'SOURCE_NAME' and 'JOB_NAME' (*TODO accept as flags*)
3. Run Loki and Grafana by `docker-compose up` (loki listens on port `3100`, grafana listens on `3000`)
3. `make build`