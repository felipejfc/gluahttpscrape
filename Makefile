PACKAGES = $(shell glide novendor)

setup-ci:
	@go get github.com/mattn/goveralls

coverage:
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg) || exit 1 &&\
		tail -n +2 coverage.out >> coverage-all.out;)

test:
	@go test
