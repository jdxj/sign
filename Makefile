# components := apiserver.out crontab.out executor.out notice.out secret.out trigger.out user.out
components := trigger.out user.out notice.out task.out app.out
images := $(subst .out,,$(components))

.PHONY: all
all: $(components)
$(components): output := _output/build
$(components):
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o $(output)/$@ cmd/$(subst .out,,$@)/*.go

.PHONY: docker
docker: $(images)
$(images): src := _output/build
$(images): des := build/docker
$(images): all
	upx -7 $(src)/$@.out
	mkdir -p $(des)/$@
	cp $(src)/$@.out $(des)/$@/$@.run
	cd $(des) && ./build.sh $@

tools := signctl.out

.PHONY: ctl
ctl: $(tools)
$(tools): output := _output/tools
$(tools):
	mkdir -p $(output)
	go build -ldflags '-s -w' -o $(output)/$@ cmd/$(subst .out,,$@)/*.go

.PHONY: clean
clean:
	rm -rf ./_output
