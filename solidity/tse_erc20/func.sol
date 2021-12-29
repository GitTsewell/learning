// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract calculation {
    function getReward(uint amount,uint multiplierX100,uint x) public pure returns(uint,uint,int) {
        uint fee = 2 * amount / 100; // 2% fee
        if(fee<2) fee = 2; //at least 2
        uint remainedAmount = amount - fee;
        uint reward = remainedAmount*multiplierX100/100;
        
        uint totalAmount = amount * 10000;
        uint totalReward = reward * x;
        int profit = int(totalReward) - int(totalAmount);
        return (totalAmount,totalReward,profit);
    }

    function mulBet() public pure returns(uint,uint,int) {
        
    }
}