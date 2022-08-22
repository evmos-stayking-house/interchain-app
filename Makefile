BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

abi: abi-gen abi-go

abi-gen:
	solc --abi $(CURDIR)/../defi-contract/contracts/token/Stayking.sol -o contracts/Stayking --base-path $(CURDIR)/../defi-contract/contracts/ --overwrite
	solc --abi $(CURDIR)/../defi-contract/contracts/token/UnbondedEvmos.sol -o contracts/UnbondedEvmos --base-path $(CURDIR)/../defi-contract/contracts/ --overwrite

abi-go:
	abigen --abi $(CURDIR)/contracts/Stayking/Stayking.abi --pkg stayking --out $(CURDIR)/abis/stayking/stayking.go
	abigen --abi $(CURDIR)/contracts/UnbondedEvmos/UnbondedEvmos.abi --pkg unbonded_evmos --out $(CURDIR)/abis/unbonded_evmos/unbonded_evmos.go

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