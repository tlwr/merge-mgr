test: *.go **/*.go
	go test -v ./...

build: *.go
	go build

install: build
	go install
