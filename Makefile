build:
	CGO_ENABLED=0 go build -o bin/rumgo

run-build: build
	bin/rumgo

br: run-build

clean:
	rm -rf bin/
