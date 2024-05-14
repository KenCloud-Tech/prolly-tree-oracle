// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


// struct ProllyConfig{
//     uint8 StrategyType;
//     uint MinNodeSize;
//     uint MaxNodeSize;
//     uint MaxPairsInNode;
//     bytes32 NodeCodec; // 假设这是合理的大小。如果不够，考虑其他方式
//     int Strategy; // 假设Strategy可以简化表示
// }

interface OracleInterface{
    function Create(string calldata dbName, string calldata primaryKey) external returns(uint reqID);
    event create(uint reqID, string dbName, string primaryKey, address owner);


    function AllowWrite(string calldata dbName, address to) external returns(uint reqID);
    // event allowWrite(uint id, address user)


    function Put(string calldata dbName, bytes calldata data) external returns(uint reqID);
    event put(uint reqID, string dbName, bytes data, address sender);


    function Get(string calldata dbName, string calldata Key, string calldata Val) external returns(uint reqID);
    event get(uint reqID, string dbName, string Key, string Val,address sender);


    // function Index(uint id, string calldata index) external;
    // event index(uint id, string index);


    // function Search(uint id, string calldata query) external view returns (uint reqID);
    // event search(uint id, string query);
}