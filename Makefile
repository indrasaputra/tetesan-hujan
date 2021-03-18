format:
	gofmt -s -w .

lint:
	golangci-lint run ./...
	
test:
	go test -v -race ./...

dep-download:
	env GO111MODULE=on go mod download

tidy:
	env GO111MODULE=on go mod tidy

vendor:
	env GO111MODULE=on go mod vendor

cover:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out 

coverhtml:
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

compile:
	env GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o tetesan-hujan-bot cmd/bot/main.go

build-docker:
	docker build -t asia.gcr.io/tetesan-hujan/tetesan-hujan-bot:v0.0.1 -f Dockerfile .

push-docker:
	docker push asia.gcr.io/tetesan-hujan/tetesan-hujan-bot:v0.0.1