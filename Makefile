BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on

abi: abi-gen abi-go abi-test

abi-test:
	solc --abi $(CURDIR)/contracts/test.sol -o contracts --overwrite
	abigen --abi $(CURDIR)/contracts/test.abi --pkg abis --out $(CURDIR)/abis/events.go

abi-gen:
	solc --abi $(CURDIR)/../defi-contract/contracts/token/Stayking.sol -o contracts/stayking --base-path $(CURDIR)/../defi-contract/contracts/ --overwrite
	solc --abi $(CURDIR)/../defi-contract/contracts/token/UnbondedEvmos.sol -o contracts/unbonded_evmos --base-path $(CURDIR)/../defi-contract/contracts/ --overwrite

abi-go:
	abigen --abi $(CURDIR)/contracts/stayking/Stayking.abi --pkg stayking --out $(CURDIR)/abis/stayking/stayking.go
	abigen --abi $(CURDIR)/contracts/unbonded_evmos/UnbondedEvmos.abi --pkg unbonded_evmos --out $(CURDIR)/abis/unbonded_evmos/unbonded_evmos.go

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