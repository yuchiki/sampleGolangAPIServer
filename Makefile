.PHONY: sampleGolangAPIServer clean

sampleGolangAPIServer:
	go build -o sampleGolangAPIServer cmd/sampleGolangAPIServer/main.go

clean:
	rm sampleGolangAPIServer
