output = _output/bin
name = sign.out
wssign = cmd/start

sign:
	mkdir -p $(output)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $(output)/$(name) $(wssign)/*.go
	upx --best $(output)/$(name)

.PHONY: clean
clean:
	find . -name "*.log" | xargs rm -f
	find . -name "*.out" | xargs rm -f
