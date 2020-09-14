build:
	@{\
			go mod vendor; \
			if [ -z "${KUBECONFIG}" ];then \
					echo "KUBECONFIG is unset, set the path to your kubeconfig"; \
			else \
					go run main.go;\
			fi	\
   }