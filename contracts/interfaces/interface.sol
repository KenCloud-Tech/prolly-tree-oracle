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
        string k;

        string str;
        int64 integer;
        uint64 uinteger;
        bytes bytess;
        bool boolean;
        string comDataType;//string,int,uint,bytes,bool

        string comIndex;
        string comOp;
    }

    function Create(string calldata dbName, string calldata primaryKey) external returns(uint reqID);
    event create(uint reqID, string dbName, string primaryKey, address owner);


    function AllowWrite(string calldata dbName, address to) external returns(uint reqID);
    // event allowWrite(uint id, address user)


    function Put(string calldata dbName, bytes calldata data) external returns(uint reqID);
    event put(uint reqID, string dbName, bytes data, address sender);


    function Get(string calldata dbName, bytes calldata recordID, string calldata callBack) external returns(uint reqID);
    event get(uint reqID, string dbName, bytes recordID, string callBack, address sender);


    function Index(string calldata dbName, string calldata Key) external returns(uint reqID);
    event index(uint reqID, string dbName, string Key, address sender);

    
    function Search(string calldata dbName, SearchController calldata Val, string calldata Method, string calldata callBack) external returns (uint reqID);
    event search(uint reqID, string dbName, SearchController Val,string Method, string callBack, address sender);
}