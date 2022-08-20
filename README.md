[Submodule] scheduled worker (with Golang) of Stayking finance

A collection of commands/services that are run to accomodate the CosmosSDK side of Stayking Finance.

[Features]
 - [x] End of epoch auto compound (withdraw rewards & delegate to the same validator
 - [x] Staking event listener
 - [x] Unbonding batch handler
 - [ ] Unbonding event listener
 - [ ] Total EVMOS asset feeder

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

[Execution]

`scheduled-worker-golang serve subscribe-delegation --from <keyname>` will subscribe to delegation events at a specific contract address.

`scheduled-worker-golang exec process-undelegation --from <keyname>` will handle batch unbonding initiation upon requests (closing position etc).

`scheduled-worker-golang keys` handles evmos key related operations.

`scheduled-worker-golang serve subscribe-undelegate --from <keyname>` listens to undelegation completion event and sends the unlocked funds to unbonding contract.

`scheduled-worker-golang serve subscribe-epoch --from <keyname>` listens to epoch end event on evmos and auto compounds the rewards distributed for the epoch.
