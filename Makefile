default:

clean:
	rm -f coverage.out
	rm -f elephant

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
