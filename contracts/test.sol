// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title test
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract test {

    uint256 delegated;
    uint256 original_unbonding;
    uint256 reserveratio_bps;

    event Stake(address indexed delegator, address indexed user, uint256 amount, uint256 share);
    event Unstake(address indexed delegator, address indexed user, uint256 amount, uint256 share);

    /**
     * @dev delegate
     * @param amount value to delegate
     */
    function delegate(uint256 amount) public {
        delegated += amount;
        emit Stake(msg.sender, msg.sender, amount, 0);
    }

    /**
     * @dev undelegate
     * @param amount value to undelegate
     */
    function undelegate(uint256 amount) public {
        delegated -= amount;
        emit Unstake(msg.sender, msg.sender, amount, 0);
    }

    /**
     * @dev Return value
     * @return value of 'number'
     */
    function retrieve() public view returns (uint256){
        return delegated;
    }

        /**
     * @dev Return value
     * @return value of 'number'
     */
    function unbonded() public view returns (uint256){
        return original_unbonding;
    }

    function supplyUnbondedToken() payable public {
        /**
            @TODO
            maybe use delegate/undelegate tx ORACLE?
         */
         original_unbonding = msg.value;
    }

    function getAccruedValue(uint256 amount) public view returns (uint256) {
        return amount / 2 * delegated;
    }
}