// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;
contract IntToRoman{

    //实现罗马数字转数整数
    function intToRoman(uint256 num) external pure returns(string memory){
        if (num < 1 || num > 3999) {
            revert("Input must be between 1 and 3999");
        }
        uint256[] memory arrInt = new uint256[](13);
        arrInt[0] = 1000;
        arrInt[1] = 900;
        arrInt[2] = 500;
        arrInt[3] = 400;
        arrInt[4] = 100;
        arrInt[5] = 90;
        arrInt[6] = 50;
        arrInt[7] = 40;
        arrInt[8] = 10;
        arrInt[9] = 9;
        arrInt[10] = 5;
        arrInt[11] = 4;
        arrInt[12] = 1;


        string[] memory symbols = new string[](13);
        symbols[0] = "M";
        symbols[1] = "CM";
        symbols[2] = "D";
        symbols[3] = "CD";
        symbols[4] = "C";
        symbols[5] = "XC";
        symbols[6] = "L";
        symbols[7] = "XL";
        symbols[8] = "X";
        symbols[9] = "IX";
        symbols[10] = "V";
        symbols[11] = "IV";
        symbols[12] = "I";

        string memory res = "";
        for(uint256 i=0;i<arrInt.length;i++){
            while (num >= arrInt[i]){
                res = string(abi.encodePacked(res,symbols[i]));
                num -= arrInt[i];
            }
            if(num ==0){
                break;
            }
        }
        return res;

    }
}