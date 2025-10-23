// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;
//合约地址:0x17E2e67c8de0cEcCF2C7cD4c8C46dD39A5EE8050
contract BeggingContract{
    // 合约所有者地址
    address private immutable owner;

    // 记录每个地址的捐赠金额
    mapping(address => uint256) private donations;

    // 捐赠事件，记录捐赠者地址和金额
    event DonationReceived(address indexed donor, uint256 amount);


    constructor(){
        owner = msg.sender;
    }

    // 仅所有者可调用的修饰符
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    // 捐赠函数，允许用户向合约发送以太币
    function donate() external payable {
        require(msg.value > 0, "Donation amount must be greater than 0");

        // 记录捐赠金额
        donations[msg.sender] += msg.value;

        // 触发捐赠事件
        emit DonationReceived(msg.sender, msg.value);
    }

    // 提取资金函数，仅所有者可调用
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");

        // 将合约中的所有资金转移给所有者
        payable(owner).transfer(balance);
    }

    // 查询指定地址的捐赠金额
    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }

    // 查询合约当前余额
    function getContractBalance() external view returns (uint256) {
        return address(this).balance;
    }

    // 查询合约所有者
    function getOwner() external view returns (address) {
        return owner;
    }
}