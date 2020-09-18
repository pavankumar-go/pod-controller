build:
	@{\
			go mod vendor; \
			if [ -z "${KUBECONFIG}" ] || [ -z "${SOURCE_NAME}" ] || [ -z "${JOB_NAME}" ];then \
					echo "set KUBECONFIG, SOURCE_NAME, JOB_NAME"; \
			else \
					go run main.go;\
			fi	\
   }

docker:
	docker-compose up --build