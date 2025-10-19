// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;
contract MergeStoredArrar{
    //合并两个有序数组 (Merge Sorted Array)
    function mergeStoredArrar(uint256[] calldata num1,uint256[] calldata num2) public pure returns(uint256[] memory){
        uint256[] memory result;
        uint256 totalLength = num1.length + num2.length;
        result = new uint256[](totalLength);

        uint256 l1 = num1.length;
        uint256 l2 = num2.length;
        uint256 i = 0;
        uint256 j = 0;
        uint256 k = 0;
        while (i< l1 && j<l2){
            if(num1[i]<num2[j]){
                result[k] = num1[i];
                i++;
            }else{
                result[k] = num2[j];
                j++;
            }
            k++;
        }

        while (i<l1){
            result[k] = num1[i];
            i++;
            k++;
        }

        while (j<l2){
            result[k] = num2[j];
            j++;
            k++;
        }

        return result;

    }
}