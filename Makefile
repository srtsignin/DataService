PACKAGE = DataService
MODULE = service
GOPATH = $(CURDIR)/.gopath
BASE = $(GOPATH)/src/$(PACKAGE)/$(MODULE)

$(BASE):
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

.PHONY: all
all: | $(BASE)
	cd $(BASE) && $(GO) build -o release/$(PACKAGE) Main.go