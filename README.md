[Submodule] scheduled worker (with Golang) of the Stayking finance

A simple process that listens a delegation/undelegation events and makes appropriate handling functions

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