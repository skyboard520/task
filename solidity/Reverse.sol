// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;
contract Reverse{
    //反转一个字符串。输入 "abcde"，输出 "edcba"
    function reverse(string memory _str) pure external returns(string memory){
        //转换成byte
        bytes memory strBytes = bytes(_str);
        uint length  = strBytes.length;
        if(length == 1){
            return _str;
        }
        for(uint256 i=0;i<length/2;i++){
            bytes1 temp = strBytes[i];
            strBytes[i] = strBytes[length -1-i];
            strBytes[length -1-i] = temp;
        }
        return string(strBytes);
    }

    function reverse1(string memory _str) pure external returns(uint256){
        //转换成byte
        bytes memory strBytes = bytes(_str);
        uint256 length  = strBytes.length;
        return length;
    }
}