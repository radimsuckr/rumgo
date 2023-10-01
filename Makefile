bin_dir = bin/
test_coverage_file = test_coverage.out

build:
	CGO_ENABLED=0 go build -o bin/rumgo

build-docker: Dockerfile
	docker build -t radimsuckr/rumgo:latest .

run-build: build
	bin/rumgo

br: run-build

clean:
	rm -rf $(bin_dir)
	rm -f $(test_coverage_file)

test:
	go test -coverprofile=$(test_coverage_file) ./...

testv:
	go test -coverprofile=$(test_coverage_file) -v ./...

testcovhtml:
	go tool cover -html=$(test_coverage_file)
