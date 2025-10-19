// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;
contract RomanToInt{
    mapping(bytes1=>uint256) private  map;
    constructor(){
         map['I'] = 1;
        map['V'] = 5;
        map['X'] = 10;
        map['L'] = 50;
        map['C'] = 100;
        map['D'] = 500;
        map['M'] = 1000;
    }
    //实现整数转罗马
    function romanToInt(string calldata s ) external view returns (uint256){

        bytes memory strBytes = bytes(s);
        uint256 length = strBytes.length;
        uint256 sum = 0;
        for(uint256 i = 0;i<length;i++){
            if(i<length-1 && map[strBytes[i]] < map[strBytes[i+1]]){
                uint256 left = map[strBytes[i+1]] - map[strBytes[i]];
                sum += left;
                i++;
            }else{
                sum += map[strBytes[i]];
            }
        }
        return sum;

    }
}