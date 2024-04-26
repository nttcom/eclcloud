fmt:
	gofmt -s -w .

fmtcheck:
	(! gofmt -s -d . | grep '^')

vet:
	go vet ./... && cd v2 && go vet ./... && cd ../v4 && go vet ./...

test:
	go test ./... -count=1 && cd v2 && go test ./... -count=1 && cd ../v4 && go test ./... -count=1

.PHONY: fmt fmtcheck vet test
