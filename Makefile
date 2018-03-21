build:
	dep ensure
	go build -ldflags="-s -w" -o bin/assetcompiler src/main.go
	cp bin/assetcompiler $(GOBIN)/assetcompiler