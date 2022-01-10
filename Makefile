all: clean build test run ## run all

clean:
	go mod tidy

run: ## run the API server
	go run cmd/family-tree/main.go

build: ## run the API server
	go build cmd/family-tree/main.go

test: ## run all test
	go test -v -cover ./...

test-p: ## run all test with coverage
	go test -p=1 -cover -covermode=count ./...

build-docker: ## build the API server as a docker image
	docker build -f ./deployments/docker/service/Dockerfile -t server .

rebuild-env: stop-env destroy-env build-env start-env

start-env:
	docker-compose -f ./deployments/docker/localstack/docker-compose-local.yml  up

stop-env:
	docker-compose -f ./deployments/docker/localstack/docker-compose-local.yml  stop

destroy-env:
	docker-compose -f ./deployments/docker/localstack/docker-compose-local.yml  rm -f

build-env:
	docker-compose -f ./deployments/docker/localstack/docker-compose-local.yml  build
