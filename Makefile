PACKAGE  = DataService
GOPATH   = $(CURDIR)/.gopath
BASE     = $(GOPATH)/src/$(PACKAGE)

$(BASE):
    @mkdir -p $(dir $@)
    @ln -sf $(CURDIR) $@

.PHONY: all
all: | $(BASE)
    cd $(BASE) && $(GO) build -o bin/$(PACKAGE) main.go