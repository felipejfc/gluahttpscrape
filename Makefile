setup-ci:
	@go get golang.org/x/tools/cmd/cover
	@go get github.com/mattn/goveralls

coverage:
	@go test -covermode=count

test:
	@go test
