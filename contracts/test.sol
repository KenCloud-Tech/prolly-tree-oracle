// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract OracleTest{

    address public oracleADD;

    function setOracle(address add) public{
        oracleADD = add;
    }

    constructor() payable {}
    receive() payable external{}


    //allow a address write
    function AllowWrite(address to, uint v) external payable returns(uint ReqID, string memory info){
        (bool success, bytes memory result) = oracleADD.call{value: v}(abi.encodeWithSignature("AllowWrite(address)", to));
        
        require(success, "Failed to call Oracle");
        (ReqID, info) = abi.decode(result, (uint256, string)); 
        return (ReqID, info);
    }

    //creat a new db
    function Create(string calldata dbName, string calldata colName, string calldata primaryKey, uint v) external payable returns(uint ReqID){
        (bool success, bytes memory result) = oracleADD.call{value: v}(abi.encodeWithSignature("Create(string,string,string)", dbName,colName,primaryKey));
    
        require(success, "Failed to call Oracle");
        ReqID = abi.decode(result, (uint256)); 
        return ReqID;
    }

    // put data to db
    function Put(string calldata dbName, string calldata colName, bytes calldata data, uint value) external payable returns(uint ReqID){
        (bool success, bytes memory result) = oracleADD.call{value: value}(abi.encodeWithSignature("Put(string,string,bytes)", dbName,colName,data));
    
        require(success, "Failed to call Oracle");
        ReqID = abi.decode(result, (uint256)); 
        return ReqID;
    }

    //get data from db
    function Get(string calldata dbName, string calldata colName, bytes calldata recordID, string calldata callBack, uint value) external payable returns(uint ReqID){
        (bool success, bytes memory result) = oracleADD.call{value: value}(abi.encodeWithSignature("Get(string,string,bytes,string)", dbName,colName,recordID,callBack));
    
        require(success, "Failed to call Oracle");
        ReqID = abi.decode(result, (uint256)); 
        return ReqID;
    }

    //creat index
    function Index(string calldata dbName, string calldata colName, string calldata Key, uint value) external payable returns(uint ReqID){
        (bool success, bytes memory result) = oracleADD.call{value: value}(abi.encodeWithSignature("Index(string,string,string)", dbName,colName,Key));
    
        require(success, "Failed to call Oracle");
        ReqID = abi.decode(result, (uint256)); 
        return ReqID;
    }
    
    //query by {equals, compare, sort, limit, skip}
    function Search(string calldata dbName, string calldata colName,bytes calldata querys, string calldata callBack, uint value) external payable returns(uint ReqID){
        (bool success, bytes memory result) = oracleADD.call{value: value}(abi.encodeWithSignature("Search(string,string,bytes,string)", dbName,colName,querys,callBack));
    
        require(success, "Failed to call Oracle");
        ReqID = abi.decode(result, (uint256)); 
        return ReqID;
    }

    function Import_from_url(string calldata dbName, string calldata colName, string calldata url, string calldata format, uint value) external payable returns(uint ReqID){
        (bool success, bytes memory result) = oracleADD.call{value: value}(abi.encodeWithSignature("Import_from_url(string,string,string,string)", dbName,colName,url,format));
    
        require(success, "Failed to call Oracle");
        ReqID = abi.decode(result, (uint256)); 
        return ReqID;
    }






    function GetCol(string calldata cid, string calldata callBack, uint value) external payable returns(uint ReqID){
        (bool success, bytes memory result) = oracleADD.call{value: value}(abi.encodeWithSignature("GetCol(string,string)", cid,callBack));
    
        require(success, "Failed to call Oracle");
        ReqID = abi.decode(result, (uint256)); 
        return ReqID;
    }

    function GetIndex(string calldata dbName, string calldata colName, string calldata callBack, uint value) external payable returns(uint ReqID){
        (bool success, bytes memory result) = oracleADD.call{value: value}(abi.encodeWithSignature("GetIndex(string,string,string)", dbName,colName,callBack));
    
        require(success, "Failed to call Oracle");
        ReqID = abi.decode(result, (uint256)); 
        return ReqID;
    }

    function getRootCid(string calldata dbName, string calldata callBack, uint value) external payable returns(uint ReqID){
        (bool success, bytes memory result) = oracleADD.call{value: value}(abi.encodeWithSignature("getRootCid(string,string)", dbName,callBack));
    
        require(success, "Failed to call Oracle");
        ReqID = abi.decode(result, (uint256)); 
        return ReqID;
    }
    
    

    // catch data and emit event
    event CatchData(bytes data);
    function CBFunc(bytes calldata data) external{
        //You can process the obtained data here.
        emit CatchData(data);
    }

}