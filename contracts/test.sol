// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.7.0 <0.9.0;

/**
 * @title Storage
 * @dev Store & retrieve value in a variable
 * @custom:dev-run-script ./scripts/deploy_with_ethers.ts
 */
contract test {

    uint256 delegated;
    uint256 original_unbonding;

    event Delegate(address indexed delegator, uint256 amount);
    event Undelegate(address indexed delegator, uint256 amount);

    /**
     * @dev delegate
     * @param amount value to delegate
     */
    function delegate(uint256 amount) public {
        delegated += amount;
        emit Delegate(msg.sender, amount);
    }

    /**
     * @dev undelegate
     * @param amount value to undelegate
     */
    function undelegate(uint256 amount) public {
        delegated -= amount;
        emit Undelegate(msg.sender, amount);
    }

    /**
     * @dev Return value
     * @return value of 'number'
     */
    function retrieve() public view returns (uint256){
        return delegated;
    }

    /**
     * @dev supplyUnbondedToken
     */
    function supplyUnbondedToken() payable public {
        /**
            @TODO
            maybe use delegate/undelegate tx ORACLE?
         */
         original_unbonding = 1234;
    }
}