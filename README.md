[Submodule] scheduled worker (with Golang) of Stayking finance

A collection of commands/services that are run to accomodate the CosmosSDK side of Stayking Finance.

[Features]
 - [x] End of epoch auto compound (withdraw rewards & delegate to the same validator
 - [x] Staking event listener
 - [x] Unbonding batch handler
 - [ ] Unbonding event listener
 - [ ] Total EVMOS asset feeder
 - [ ] Storage for bookkeeping (last undelegation tx, handled blocks etc)

[Terminologies]

Delegating = Staking

Unbonding = Unstaking

[Handlers]

 - Delegation
   - grabs the from address and amount
   - Converts the incoming wrapped EVMOS to EVMOS
   - Makes equivalent (?) amount delegation
 - Undelegation
   - grabs the from address and amount
   - Appends the entry to a waiting list
   - flushes all waiting list to start an unbonding(s) to get a sufficient amount of EVMOS every few days.
   - When an unbonding is complete, converts the equivalent amount to wrapped EVMOS and calls the staking contract for repayment.

[Processes]

 - Total EVMOS asset value query & update to contract to get stEVMOS:EVMOS exchange rate
   - Should happen at least once a day
 - Undelegation tracking & processing
   - Need to keep track of current unbonding waitlist to be processed in the next batch
   - Batch should happen once every ~2 days or so to avoid hitting the max unbonding entries
     - EVMOS unbonding period = 14 days
     - Max unbonding entries = 7

[Build]

 - Install Golang
 - make clean && make build
 - build/scheduled-worker-golang

 - Initialize `~/.evmosd` environment through either evmosd command or manual mkdir.
 - configure keyring to `test` through `evmosd config keyring-backend test` to avoid having to enter password for each tx
 - Add a key that has some amount of funds (should be small just enough to make txs)

[Execution]

`scheduled-worker-golang serve subscribe --from <keyname> --cont-addr <stayking contract addr> --validator <validator addr>` will subscribe to delegation, unbonding, epoch end events.

