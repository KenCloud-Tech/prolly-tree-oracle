// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OracleInterface} from "./interfaces/interface.sol";
import {IOracle} from "./I_oracle.sol";
import {Permission} from "./interfaces/permission.sol";
import {util} from "./utils.sol";

contract Oracle is IOracle, OracleInterface, util {
    //allow a address write
    function AllowWrite(address to) external override{
        require( !isEmptyString(myDbName[msg.sender]), "You has not create a db.");
        string memory dbName = myDbName[msg.sender];
        require(msg.sender == dbOwner[dbName], "Only the db owner can call this function");
        uint reqID = CurrentReqID;
        CurrentReqID++;
        permission[to][dbName] = Permission(true, true, true);
        emit ReqState(reqID, msg.sender, true, "Permission granted successfully");
        reqStatement[reqID] = true;
    }

    //creat a new collection
    function Create(string calldata dbName, string calldata colName, string calldata primaryKey)external override {
        require(dbOwner[dbName]==address(0) || dbOwner[dbName]==msg.sender || permission[msg.sender][dbName].allowWrite == true,  "You do not have permission to write");
        require(cols[dbName][colName] == false, "This collection had been created.");
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit create(reqID, dbName, colName, primaryKey, msg.sender);
    }

    function CreatRsp(uint reqID, bool statement,string calldata dbName, string calldata colName, address sender, string calldata info) onlyOracleOwner external {
        if (statement == true) {
            if (dbOwner[dbName] ==address(0)){
                dbOwner[dbName] = sender;
                myDbName[sender] = dbName;
            }
            cols[dbName][colName] = true;
            permission[dbOwner[dbName]][dbName] = Permission(true, true, true); //dbOwner
            permission[sender][dbName] = Permission(true, true, true); //clollection creater
            emit ReqState(reqID, sender, true, "Collection creat success.");
            reqStatement[reqID] = true;
        } else {
            emit ReqState(reqID, sender, false, info);
            reqStatement[reqID] = false;
        }
    }

    // put data to collection
    function Put(string calldata dbName, string calldata colName, bytes calldata data) external override colIsExsist(dbName, colName) allowWrite(dbName) {
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit put(reqID, dbName, colName, data, msg.sender);
    }

    function PutRsp(uint reqID, bool statement, address sender, string calldata info) onlyOracleOwner external {
        if (statement == true) {
            emit ReqState(reqID, sender, true, "Put data success.");
            reqStatement[reqID] = true;
        } else {
            emit ReqState(reqID, sender, false, info);
            reqStatement[reqID] = false;
        }
    }

    //get data from collection
    function Get(string calldata dbName, string calldata colName, bytes calldata recordID, string calldata callBack) external override colIsExsist(dbName, colName) allowQuery(dbName) {
        emit get(CurrentReqID++, dbName, colName, recordID, callBack, msg.sender);
    }

    function GetRsp(uint reqID, bool statement, bytes calldata data, string calldata callBack, address sender, string calldata info) onlyOracleOwner external {
        if (statement == true) {
            (bool OK,) = sender.call(abi.encodeWithSignature(callBack, data));
            if (OK) {
                emit ReqState(reqID, sender, true, "Get data success.");
                reqStatement[reqID] = true;
                return;
            }
        }

        emit ReqState(reqID, sender, false, info);
        reqStatement[reqID] = false;
    }


    //creat index
    function Index(string calldata dbName, string calldata colName, string calldata idx) external colIsExsist(dbName, colName) allowWrite(dbName) override {
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit index(reqID, dbName, colName, idx, msg.sender);
    }

    function IndexRsp(uint reqID, bool statement, address sender, string calldata info) onlyOracleOwner external {
        if (statement == true) {
            emit ReqState(reqID, sender, true, "Creat index success.");
            reqStatement[reqID] = true;
        } else {
            emit ReqState(reqID, sender, false, info);
            reqStatement[reqID] = false;
        }
    }



    //query by {equals, compare, sort, limit, skip}
    function Search(string calldata dbName, string calldata colName, bytes calldata querys, string calldata callBack) external colIsExsist(dbName, colName) allowQuery(dbName) override {
        emit search(CurrentReqID++, dbName, colName, querys, callBack, msg.sender);
    }

    function SearchRsp(uint reqID, bool statement, bytes calldata data, string calldata callBack, address sender, string calldata info) onlyOracleOwner external {
        if (statement == true) {
            (bool OK,) = sender.call(abi.encodeWithSignature(callBack, data));
            if (OK) {
                emit ReqState(reqID, sender, true, "Get data success.");
                reqStatement[reqID] = true;
                return;
            }
        }

        emit ReqState(reqID, sender, false, info);
        reqStatement[reqID] = false;
    }

}