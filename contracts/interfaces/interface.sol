// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


// struct ProllyConfig{
//     uint8 StrategyType;
//     uint MinNodeSize;
//     uint MaxNodeSize;
//     uint MaxPairsInNode;
//     bytes32 NodeCodec; 
//     int Strategy;
// }

interface OracleInterface{

    function Create(string calldata cid, string calldata colName, string calldata primaryKey) external;
    event create(uint reqID,string cid, string colName, string primaryKey, address owner);


    function AllowWrite(address to) external;
    // event allowWrite(uint id, address user)


    function Put(string calldata cid, string calldata colName, bytes calldata data) external;
    event put(uint reqID, string cid, string colName, bytes data, address sender);


    function Get(string calldata cid, string calldata colName, bytes calldata recordID, string calldata callBack) external;
    event get(uint reqID, string cid, string colName, bytes recordID, string callBack, address sender);


    function Index(string calldata cid, string calldata colName, string calldata idx) external;
    event index(uint reqID, string cid, string colName, string Key, address sender);


    function Search(string calldata cid, string calldata colName, bytes calldata querys, string calldata callBack) external;
    event search(uint reqID, string cid, string colName, bytes querys, string callBack, address sender);
}