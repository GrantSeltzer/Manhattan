VERSION=0.0.1
BINARY=Manhattan
binary:
	go build -ldflags "-X main.version=${VERSION}"
clean:
	rm -f Manhattan
test:
	cd tests && go test
