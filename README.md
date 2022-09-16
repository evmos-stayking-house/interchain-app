##  [Submodule] scheduled worker (with Golang) of StayKing House

A collection of commands/services that are run to accommodate the CosmosSDK side of StayKing House

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
make clean && make build && build/scheduled-worker-golang serve subscribe --from bob --cont-addr 0xd8A9159c111D0597AD1b475b8d7e5A217a1d1d05 --uevmos-cont-addr 0x8b9d5A75328b5F3167b04B42AD00092E7d6c485c --validator evmosvaloper1qvc6jej73armfs5fadn9lprx768f46d9wpd7d7 --broadcast-mode async  --eth-endpoint http://eth.bd.evmos.dev:8545 --node http://bd-evmos-testnet-state-sync-node-01.bdnodes.net:26657
```
