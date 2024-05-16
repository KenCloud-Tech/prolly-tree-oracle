// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OracleInterface} from "./interfaces/interface.sol";
import {IOracle} from "./I_oracle.sol";
import {Permission} from "./interfaces/permission.sol";

contract Oracle is IOracle, OracleInterface {
    //allow a address write
    function AllowWrite(string calldata dbName, address to) external override onlyDbOwner(dbName) {
        uint reqID = CurrentReqID;
        CurrentReqID++;
        permission[to][dbName] = Permission(true, true, true);
        emit ReqState(reqID, msg.sender, true, "Permission granted successfully");
        reqStatement[reqID] = true;
    }

    //creat a new db
    function Create(string calldata dbName, string calldata primaryKey) external override {
        require(dbOwner[dbName] == address(0), "This db had been created.");
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit create(reqID, dbName, primaryKey, msg.sender);
    }

    function CreatRsp(uint reqID, bool statement, string calldata dbName, address sender, string calldata info) onlyOracleOwner external {
        if (statement == true) {
            dbOwner[dbName] = sender;
            permission[sender][dbName] = Permission(true, true, true);
            emit ReqState(reqID, sender, true, "Db creat success.");
            reqStatement[reqID] = true;
        } else {
            emit ReqState(reqID, sender, false, info);
            reqStatement[reqID] = false;
        }
    }

    // put data to db
    function Put(string calldata dbName, bytes calldata data) external override allowWrite(dbName) {
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit put(reqID, dbName, data, msg.sender);
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

    //get data from db
    function Get(string calldata dbName, bytes calldata recordID, string calldata callBack) external override allowQuery(dbName) {
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit get(reqID, dbName, recordID, callBack, msg.sender);
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
    function Index(string calldata dbName, string calldata Key) external allowWrite(dbName) override {
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit index(reqID, dbName, Key, msg.sender);
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
    function Search(string calldata dbName, SearchController calldata Val, string calldata Method, string calldata callBack) external allowQuery(dbName) override {
        bytes32 method = keccak256(abi.encodePacked(Method));
        require(method == equal || method == compare || method == sort || method == limit || method == skip, "The search method is wrong, only supports: equal,compare,sort,limit and skip.");
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit search(reqID, dbName, Val, Method, callBack, msg.sender);
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