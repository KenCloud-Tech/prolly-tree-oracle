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


    function AllowWrite(address to) external payable returns(uint ReqID, string memory info);
    // event allowWrite(uint id, address user)


    function Create(string calldata dbName, string calldata colName, string calldata primaryKey) external payable returns(uint ReqID);
    event create(uint reqID,string dbName, string colName, string primaryKey, address owner);


    function Put(string calldata dbName, string calldata colName, bytes calldata data) external payable returns(uint ReqID);
    event put(uint reqID, string dbName, string colName, bytes data, address sender);


    function Get(string calldata dbName, string calldata colName, bytes calldata recordID, string calldata callBack) external payable returns(uint ReqID);
    event get(uint reqID, string dbName, string colName, bytes recordID, string callBack, address sender);


    function Index(string calldata dbName, string calldata colName, string calldata idx) external payable returns(uint ReqID);
    event index(uint reqID, string dbName, string colName, string Key, address sender);


    function Search(string calldata dbName, string calldata colName, bytes calldata querys, string calldata callBack) external payable returns(uint ReqID);
    event search(uint reqID, string dbName, string colName, bytes querys, string callBack, address sender);

    function Import_from_url(string calldata dbName, string calldata colName, string calldata url, string calldata format) external payable returns(uint ReqID); //format={ndjson or csv}
    event import_from_url(uint reqID, string dbName, string colName, string url, string format, address sender);
}