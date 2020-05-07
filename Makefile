test:
	env go test -count=1 ./...

test-all:
	env go test -count=1 ./... -tags=integration 

.PHONY: test