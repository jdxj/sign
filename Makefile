localName = sign.out
linuxName = sign_linux.out
macName = sign_mac.out

local: clean
	go build -ldflags '-s -w' -o $(localName) *.go
linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $(linuxName) *.go
	upx --best $(linuxName)
mac: clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o $(macName) *.go
	upx --best $(macName)
clean:
	find . -name "*.log" | xargs rm -f
	find . -name "*.out" | xargs rm -f
