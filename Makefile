# components := apiserver.out crontab.out executor.out notice.out secret.out trigger.out user.out
components := trigger.out user.out notice.out task.out app.out
images := $(subst .out,,$(components))

.PHONY: all
all: $(components)
$(components): output := build/output
$(components):
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $(output)/$@ cmd/$(subst .out,,$@)/*.go

.PHONY: docker
docker: $(images)
$(images): src := build/output
$(images): des := build/docker
$(images): all
	upx -7 $(src)/$@.out
	mkdir -p $(des)/$@
	cp $(src)/$@.out $(des)/$@/$@.run
	cd $(des) && ./build.sh $@

tools := signctl.out

.PHONY: ctl
ctl: $(tools)
$(tools): output := build/tools
$(tools):
	mkdir -p $(output)
	go build -ldflags '-s -w' -o $(output)/$@ cmd/$(subst .out,,$@)/*.go

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -rf build/output
	rm -rf build/tools
