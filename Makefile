output := _output/bin

sign: name := sign.out
sign: package := cmd/sign
sign:
	mkdir -p $(output)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $(output)/$(name) $(package)/*.go
	upx --best $(output)/$(name)

api: name := apiserver.out
api: package := cmd/apiserver
api:
	mkdir -p $(output)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $(output)/$(name) $(package)/*.go
	upx --best $(output)/$(name)

.PHONY: clean
clean:
	find . -name "*.log" | xargs rm -f
	find . -name "*.out" | xargs rm -f
