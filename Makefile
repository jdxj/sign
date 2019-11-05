local:
	go build -o sign.out *.go
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o sign.out *.go
mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o sign.out *.go
clean:
	rm -rvf *.out *.log
