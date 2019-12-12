fmt:
	gofmt -s -w .

fmtcheck:
	(! gofmt -s -d . | grep '^')

test: fmtcheck
	go test ./... -count=1

.PHONY: fmt fmtcheck test
