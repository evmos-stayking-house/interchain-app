BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

abi:
	abigen --abi $(CURDIR)/contracts/test.abi --pkg events --out $(CURDIR)/events/events.go

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/
build-linux:
	GOOS=linux GOARCH=amd64 LEDGER_ENABLED=false $(MAKE) build

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

clean:
	rm -rf \
    $(BUILDDIR)/