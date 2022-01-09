.PHONY: sampleGolangAPIServer serve browse test clean

sampleGolangAPIServer:
	go build -o sampleGolangAPIServer -ldflags="-X main.VERSION=untagged" cmd/sampleGolangAPIServer/main.go

serve:
	./sampleGolangAPIServer

browse:
	curl localhost:8080/hello

test:
	go test -v ./...

clean:
	rm sampleGolangAPIServer
