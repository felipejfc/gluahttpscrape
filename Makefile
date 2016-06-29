setup-ci:
	@sudo add-apt-repository -y ppa:masterminds/glide && sudo apt-get update
	@sudo apt-get install -y glide
	@go get golang.org/x/tools/cmd/cover
	@go get github.com/mattn/goveralls
	@glide install

coverage:
	@go test -covermode=count

test:
	@go test

deps:
	@glide install
