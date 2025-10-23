// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;
import "./IERC20.sol";
contract ERC20 is IERC20{
    mapping(address => uint256)  balances;
    mapping(address => mapping(address => uint256)) allowances;
    uint public totalSupply;//代币总供给
    string public name; // 名称
    string public symbol;//符号
    uint public decimals = 18;//小数点位数

    //在合约部署的时候实现合约名称和符号
    constructor(string memory _name, string memory _symbol){
        name = _name;
        symbol = _symbol;
    }

    function balanceOf(address _owner) public view returns (uint256 balance) {
        return balances[_owner];
    }

    function allowance(address _owner, address _spender) public view returns (uint256 remaining) {
        return allowances[_owner][_spender];
    }

    /**
     * @dev 转账 `amount` 单位代币，从调用者账户到另一账户 `to`.
     *
     * 如果成功，返回 `true`.
     *
     * 释放 {Transfer} 事件.
     */
    function transfer(address to, uint256 amount) external returns (bool){
        balances[msg.sender] -= amount;
        balances[to] += amount; // 更新余额]
        emit Transfer(msg.sender,to,amount);
        return true;
    }

    /**
     * @dev 调用者账户给`spender`账户授权 `amount`数量代币。
     *
     * 如果成功，返回 `true`.
     *
     * 释放 {Approval} 事件.
     */
    function approve(address spender, uint256 amount) external  returns (bool){
        allowances[msg.sender][spender] = amount;
        emit Approval(msg.sender,spender,amount);
        return true;
    }

    /**
     * @dev 通过授权机制，从`from`账户向`to`账户转账`amount`数量代币。转账的部分会从调用者的`allowance`中扣除。
     *
     * 如果成功，返回 `true`.
     *
     * 释放 {Transfer} 事件.
     */
    function transferFrom(address from,address to,uint256 amount) external returns (bool){
        require(balances[from] >= amount,"not ");
        require(allowances[from][msg.sender] >= amount,"not ");
        balances[from] -= amount;
        balances[to] += amount;
        allowances[from][msg.sender] -= amount;
        emit Transfer(from,to,amount);
        return true;
    }

    // @dev 铸造代币，从 `0` 地址转账给 调用者地址
    function mint(uint amount) external {
        balances[msg.sender] += amount;
        totalSupply += amount;
        emit Transfer(address(0), msg.sender, amount);
    }

    // @dev 销毁代币，从 调用者地址 转账给  `0` 地址
    function burn(uint amount) external {
        balances[msg.sender] -= amount;
        totalSupply -= amount;
        emit Transfer(msg.sender, address(0), amount);
    }


}