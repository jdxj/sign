build.%: output := _output/build
build.%:
	mkdir -p $(output)
	go build -o $(output)/$*.out cmd/$*/*.go

cross.%: output := _output/cross
cross.%:
	mkdir -p $(output)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $(output)/$*.out cmd/$*/*.go
	upx --best $(output)/$*.out

.PHONY: clean
clean:
	find . -name "*.log" | xargs rm -f
	find . -name "*.out" | xargs rm -f
