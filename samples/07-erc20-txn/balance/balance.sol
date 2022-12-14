// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

abstract contract ERC20TokenBalance {
    function balanceOf(address account) public view virtual returns (uint256);
}