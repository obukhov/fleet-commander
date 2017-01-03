NAME=fleetcmd
SOURCE=cmd/fleetcmd/main.go

# dependencies that are used by the build&test process, these need to be installed in the
DEPEND=github.com/jstemmer/go-junit-report github.com/Masterminds/glide
DATE=$(shell date '+%F %T')
COMMIT?=$(shell git rev-parse HEAD)

# produce a version string that is embedded into the binary that captures the date and the commit we're building
VERSION=$(COMMIT)-$(DATE)
LDFLAGS=-ldflags "-X 'main.VERSION=$(VERSION)'"
GOBUILD=go build $(LDFLAGS)
GOTEST=go test -v

all: clean depend test build
build:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o build/linux_amd64/$(NAME) $(SOURCE)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o build/darwin_amd64/$(NAME) $(SOURCE)
test:
	$(GOTEST) $(shell glide novendor)
test_report:
	mkdir build/
	$(GOTEST) $(shell glide novendor) > build/report.out
	cat build/report.out | go-junit-report --set-exit-code=true > build/report.xml
clean:
	rm -rf build/
# installing build dependencies. You will need to run this once manually when you clone the repo
depend:
	go get -u -v $(DEPEND)
	glide install
install:
	go install $(SOURCE)
