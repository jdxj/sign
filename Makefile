components := apiserver crontab executor notice secret trigger user
.PHONY: all
all: $(components)
$(components): output := _output/build
$(components):
	mkdir -p $(output)
	go build -o $(output)/$@.out cmd/$@/*.go

build.%: output := _output/build
build.%:
	mkdir -p $(output)
	go build -o $(output)/$*.out cmd/$*/*.go

cross.%: output := _output/cross
cross.%:
	mkdir -p $(output)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $(output)/$*.out cmd/$*/*.go
	upx --best $(output)/$*.out

containers := apiserver crontab executor notice secret trigger user
.PHONY: docker
docker: $(containers)
$(containers): output := build/docker
$(containers):
	mkdir -p $(output)/$@
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $(output)/$@/$@.out cmd/$@/*.go
	upx --best $(output)/$@/$@.out

.PHONY: clean
clean:
	rm -rf ./_output
