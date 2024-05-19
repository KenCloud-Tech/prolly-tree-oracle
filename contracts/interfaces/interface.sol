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

    struct SearchController{
        string k; // the key for search
        // the value for search , just use one
        string str;
        int64 integer;
        uint64 uinteger;
        bytes bytess;
        bool boolean;

        string DataType; // type of value string,int,uint,bytes or bool,
        string comOp; // compare optiong: "GreaterThan" or "LessThan"
    }

    function Create(string calldata dbName, string calldata primaryKey) external;
    event create(uint reqID, string dbName, string primaryKey, address owner);


    function AllowWrite(string calldata dbName, address to) external;
    // event allowWrite(uint id, address user)


    function Put(string calldata dbName, bytes calldata data) external;
    event put(uint reqID, string dbName, bytes data, address sender);


    function Get(string calldata dbName, bytes calldata recordID, string calldata callBack) external;
    event get(uint reqID, string dbName, bytes recordID, string callBack, address sender);


    function Index(string calldata dbName, string calldata Key) external;
    event index(uint reqID, string dbName, string Key, address sender);


    function Search(string calldata dbName, SearchController calldata Val, string calldata Method, string calldata callBack) external;
    event search(uint reqID, string dbName, SearchController Val,string Method, string callBack, address sender);
}