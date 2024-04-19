t="coverage.txt"

test:
	go test ./... -cover -v

coverage:
	go test -coverprofile=$t ./... && go tool cover -html=$t && unlink $t

