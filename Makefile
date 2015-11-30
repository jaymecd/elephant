default:

clean:
	rm -f coverage.out
	rm -f elephant

# Checks project and source code if everything is according to standard
.PHONY: check
check:
	gofmt -l . | read && echo "Code differs from gofmt's style" && exit 1 || true
	go vet ./...

test:
	go test ./...
#	go tool cover -func=coverage.out
#	go tool cover -html=coverage.out

build:
	go build -i -o elephant \
		-ldflags "-s \
			-X main.version=1.0.2 \
			-X main.buildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` \
			-X main.gitHash=`git rev-parse HEAD`" \
		main.go
