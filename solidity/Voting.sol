// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract Voting{
    //一个mapping来存储候选人的得票数
    mapping ( address => uint256 ) private  votesReceived;

    //记录候选人
    address[] public candidateList;
    //一个vote函数，允许用户投票给某个候选人
    function vote(address _address) external {
        votesReceived[_address] += 1;
        candidateList.push(_address);

    }

    //一个getVotes函数，返回某个候选人的得票数
    function getVotes(address _address) external view returns (uint256){
        return votesReceived[_address];
    }

    //一个resetVotes函数，重置所有候选人的得票数
    function resetVotes() external {
        for(uint256 i =0;i<candidateList.length;i++){
            address candidate = candidateList[i];
            votesReceived[candidate] = 0;
        }

        delete candidateList;
    }

    function size() external view returns (uint256){
        return candidateList.length;
    }


}