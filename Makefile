local:
	go build -ldflags '-s -w' -o sign.out *.go
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o sign.out *.go
	upx --best sign.out
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o sign.out *.go
	upx --best sign.out
clean:
	rm -rvf *.out *.log
