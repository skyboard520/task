// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;
contract BinarySearch{
    //二分查找 (Binary Search)
    function binarySearch(uint256[] calldata nums,uint256 target) external pure returns(int256){
        uint256 left = 0;
        uint256 right = nums.length-1;
        while (left <= right){
            uint256 mid = left + (right -left)/2;
            if(nums[mid] == target){
                return int256(mid);
            }else if(nums[mid] < target){
                left = mid +1;
            }else{
                right = mid -1;
            }
        }
        return -1;
    }
}