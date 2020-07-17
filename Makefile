build:
	env GOOS=linux CGO_ENABLED=0 go build ${LDFLAGS} -a -installsuffix cgo -o bin/questionmate world/http/main.go

deploy: build
	scp bin/questionmate 95.217.222.60:~/questionmate/bin
	scp -r config/* 95.217.222.60:~/questionmate/config

test:
	env go test -count=1 ./...

test-all:
	env go test -count=1 ./... -tags=integration 

.PHONY: test