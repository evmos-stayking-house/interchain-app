##  [Submodule] scheduled worker (with Golang) of StayKing House

A collection of commands/services that are run to accommodate the CosmosSDK side of StayKing House

## Delegator Address
- [Bech32]
  - evmos1ln2pgz7wupxzuprg6q99kl0a9r05z7zwg9fhs3
- [Hex]
  - 0xFCD4140BCeE04C2E0468d00a5b7DFd28Df41784e
- Testnet MintScan Explorer for StayKing Delegator
  - [Check Delegating/UnDelegating Transactions](https://testnet.mintscan.io/evmos-testnet/account/evmos1ln2pgz7wupxzuprg6q99kl0a9r05z7zwg9fhs3)

## Features

 - End of epoch auto compound (withdraw rewards & delegate to the same validator)
 - Distribute lender rewards & etc
 - Staking event listener
 - Unbonding event listener & unbonded coins handler (debug)
 - Total EVMOS asset feeder
 - Storage for bookkeeping (last undelegation tx, handled blocks etc)
 - Unbonding batch handler (blocked by storage)

## Terminologies

Delegating = Staking

Unbonding = Unstaking

## Handlers

 - Delegation
   - grabs the from address and amount
   - Converts the incoming wrapped EVMOS to EVMOS
   - Makes equivalent (?) amount delegation
 - Undelegation
   - grabs the from address and amount
   - Appends the entry to a waiting list
   - flushes all waiting list to start an unbonding(s) to get a sufficient amount of EVMOS every few days.
   - When an unbonding is complete, converts the equivalent amount to wrapped EVMOS and calls the staking contract for repayment.

## Processes

 - Total EVMOS asset value query & update to contract to get stEVMOS:EVMOS exchange rate
   - Should happen at least once a day
 - Undelegation tracking & processing
   - Need to keep track of current unbonding waitlist to be processed in the next batch
   - Batch should happen once every ~2 days or so to avoid hitting the max unbonding entries
     - EVMOS unbonding period = 14 days
     - Max unbonding entries = 7

## Build

 - Install Golang
 - make clean && make build
 - build/scheduled-worker-golang

 - Initialize `~/.evmosd` environment through either evmosd command or manual mkdir.
 - configure keyring to `test` through `evmosd config keyring-backend test` to avoid having to enter password for each tx
 - Add a key that has some amount of funds (should be small just enough to make txs)


## Execution

`scheduled-worker-golang serve subscribe --from <keyname> --cont-addr <stayking contract addr> --validator <validator addr>` will subscribe to delegation, unbonding, epoch end events.

example (testnet):
```bash
make clean && make build && build/scheduled-worker-golang serve subscribe --from bob --cont-addr {Deployed StayKing Contarct Address} --uevmos-cont-addr {Deployed uEVMOS Contarct Address} --validator evmosvaloper1qvc6jej73armfs5fadn9lprx768f46d9wpd7d7 --broadcast-mode async  --eth-endpoint http://eth.bd.evmos.dev:8545 --node http://bd-evmos-testnet-state-sync-node-01.bdnodes.net:26657
```
