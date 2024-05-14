// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OracleInterface} from "./interfaces/interface.sol";
import {IOracle} from "./I_oracle.sol";
import {Permission} from "./interfaces/permission.sol";

contract Oracle is IOracle,OracleInterface{
    //allow a address write
    function AllowWrite(string calldata dbName, address to) external override onlyDbOwner(dbName) returns(uint reqID){
        reqID = CurrentReqID;
        permission[to][dbName] = Permission(true, true, true);
        reqStatement[msg.sender][reqID] = ReqController(reqID, "", true);
        CurrentReqID++;
        return reqID;
    }

    //creat a new db
    function Create(string calldata dbName, string calldata primaryKey) external override returns(uint reqID){
        require(dbOwner[dbName] == msg.sender || dbOwner[dbName] == address(0), "This db had been created by other user.");
        reqID = CurrentReqID;
        CurrentReqID++;
        if (dbOwner[dbName] == msg.sender){
            reqStatement[msg.sender][reqID] = ReqController(reqID, "", true);
            return reqID;
        }
        dbOwner[dbName] = msg.sender;
        reqStatement[msg.sender][reqID] = ReqController(reqID, "", false);
        emit create(reqID, dbName, primaryKey, msg.sender);
        return reqID;
    }
    function CreatRsp(uint reqID, bool statement, string calldata dbName, address owner) onlyOracleOwner external{
        reqStatement[owner][reqID].statement = statement;
        if (statement == true){
            permission[owner][dbName] = Permission(true, true, true);
        }
    }

    // put data to db
    function Put(string calldata dbName, bytes calldata data) external override allowWrite(dbName) returns (uint reqID){
        reqID = CurrentReqID;
        CurrentReqID++;
        reqStatement[msg.sender][reqID] = ReqController(reqID, data, false);
        emit put(reqID, dbName, data, msg.sender); 
        return reqID;
    }
    function PutRsp(uint reqID, bool statement, address sender) onlyOracleOwner external{
        reqStatement[sender][reqID].statement = statement;
    }

    //get data from db
    function Get(string calldata dbName, string calldata Key, string calldata Val) external override allowQuery(dbName) returns(uint reqID){
        reqID = CurrentReqID;
        CurrentReqID++;
        reqStatement[msg.sender][reqID] = ReqController(reqID, "", false);
        emit get(reqID, dbName, Key, Val, msg.sender); 
        return reqID;
    }
    function GetRsp(uint reqID, bool statement, bytes calldata data, address sender) onlyOracleOwner external{
        reqStatement[sender][reqID].statement = statement;
        if (statement == true){
            reqStatement[sender][reqID].data = data;
            reqStatement[sender][reqID].statement = statement;
        }
    }

}