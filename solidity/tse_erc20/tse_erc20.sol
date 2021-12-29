// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TSEToken is ERC20 {
    constructor(uint256 initialSupply) ERC20("Tsewell", "TSE") {
        _mint(msg.sender, initialSupply);
    }

    function decimals() public override pure returns(uint8) {
        return 2;
    }
}