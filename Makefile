build:
	@go build -o ./bin/qrcode-generator

run: build
	@./bin/qrcode-generator

test:
	@go test -v ./...

clean:
	@rm -rf bin
	@rm -rf *.png