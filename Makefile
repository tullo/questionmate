build:
	env GOOS=linux CGO_ENABLED=0 go build ${LDFLAGS} -a -installsuffix cgo -o bin/questionmate world/http/main.go

deploy: build
	ssh 95.217.222.60 "pkill questionmate"
	scp bin/questionmate 95.217.222.60:~/questionmate/bin
	scp -r config/* 95.217.222.60:~/questionmate/config
	ssh 95.217.222.60 "nohup /home/ralf/questionmate/bin/questionmate -filename=/home/ralf/questionmate/config/legacylab-short &"

test:
	env go test -count=1 ./...

test-all:
	env go test -count=1 ./... -tags=integration

clean:
	rm -rf bin

.PHONY: test