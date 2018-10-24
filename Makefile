PACKAGE = DataService
MODULE = service
BASE = $(GOPATH)/src/$(PACKAGE)

$(BASE):
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

.PHONY: all
all: | $(BASE)
	cd $(BASE)/$(MODULE) && go build -o /release/$(PACKAGE) Main.go