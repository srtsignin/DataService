PACKAGE = DataService
MODULE = service
BASE = $(GOPATH)/src/$(PACKAGE)/$(MODULE)

$(BASE):
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

.PHONY: all
all: | $(BASE)
	cd $(BASE) && go build -o release/$(PACKAGE) Main.go