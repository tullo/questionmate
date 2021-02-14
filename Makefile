SHELL = /bin/bash -o pipefail
build:
	go build -o bin/questionmate world/http/main.go

build-linux: GOOS=linux, CGO_ENABLED=0
build-linux:
	@go build ${LDFLAGS} -a -installsuffix cgo -o bin/questionmate world/http/main.go

clean:
	@rm -rf bin

deploy: build
	@ssh 95.217.222.60 "pkill questionmate"
	@scp bin/questionmate 95.217.222.60:~/questionmate/bin
	@scp -r config/* 95.217.222.60:~/questionmate/config
	@ssh 95.217.222.60 "sh -c 'nohup /home/ralf/questionmate/bin/questionmate -directory=/home/ralf/questionmate/config/coma > /dev/null 2>&1 &'"

console:
	@go run world/console/main.go

run:
	go run ./world/http/main.go -directory=${PWD}/config > /dev/null 2>&1

test: export SRC_ROOT=${PWD}
test:
	@go test -v -count=1 -run . ./...

test-all: export SRC_ROOT=${PWD}
test-all:
	@go test -count=1 ./... -tags=integration
