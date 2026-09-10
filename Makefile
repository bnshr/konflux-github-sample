.PHONY: run test docker-build docker-run tidy

run:
	go run .

test:
	go test ./...

tidy:
	go mod tidy

docker-build:
	docker build -t konflux-github-sample:local .

docker-run: docker-build
	docker run --rm -p 8080:8080 konflux-github-sample:local
