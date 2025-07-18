PROJECT?=github.com/donskova1ex/ref_system
API_NAME?=ref_system
API_VERSION?=0.0.1
API_CONTAINER_NAME?=docker.io/donskova1ex/${API_NAME}


clean_api:
	rm -rf bin/ref_system

api_docker_build:
	docker build --no-cache -t ${API_CONTAINER_NAME}:${API_VERSION} -t ${API_CONTAINER_NAME}:latest -f dockerfile.api .

api_local_run:
	go run ./cmd/api/ref_system_api.go

