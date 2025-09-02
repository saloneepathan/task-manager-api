run:
	go run main.go

test:
	go test ./...

docker-build:
	docker build -t task-manager-api .

docker-run:
	docker run -p 8080:8080 task-manager-api
