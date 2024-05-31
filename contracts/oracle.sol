// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OracleInterface} from "./interfaces/interface.sol";
import {IOracle} from "./I_oracle.sol";
import {Permission} from "./interfaces/permission.sol";
import {util} from "./utils.sol";

contract Oracle is IOracle, OracleInterface, util {
    //allow a address write
    function AllowWrite(address to) external override{
        require( !isEmptyString(getcid[msg.sender]), "You has not create a db.");
        string memory cid = getcid[msg.sender];
        require(msg.sender == dbOwner[cid], "Only the db owner can call this function");
        uint reqID = CurrentReqID;
        CurrentReqID++;
        permission[to][msg.sender] = Permission(true, true, true);
        emit ReqState(reqID, msg.sender, true, "Permission granted successfully");
        reqStatement[reqID] = true;
    }

    //creat a new collection
    function Create(string calldata cid, string calldata colName, string calldata primaryKey) allowWrite(cid) external override {
        require(cols[dbOwner[cid]][colName] == false, "This collection had been created.");
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit create(reqID, cid, colName, primaryKey, msg.sender);
    }

    function CreatRsp(uint reqID, bool statement, string calldata newcid, string calldata oldcid, string calldata colName, address sender, string calldata info) onlyOracleOwner external {
        if (statement == true) {
            address owner;
            if (isEmptyString(oldcid)){
                owner = sender;
            } else{
                owner = dbOwner[oldcid];
            }
            dbOwner[newcid] = owner;
            getcid[owner] = newcid;
            cols[owner][colName] = true;
            permission[owner][owner] = Permission(true, true, true); //dbOwner
            permission[sender][owner] = Permission(true, true, true); //clollection creater
            emit ReqState(reqID, sender, true, "Db creat success.");
            emit NewCid(sender,newcid);
            reqStatement[reqID] = true;
        } else {
            emit ReqState(reqID, sender, false, info);
            reqStatement[reqID] = false;
        }
    }

    // put data to collection
    function Put(string calldata cid, string calldata colName, bytes calldata data) external override colIsExsist(cid, colName) allowWrite(cid) {
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit put(reqID, cid, colName, data, msg.sender);
    }

    function PutRsp(uint reqID, bool statement, string calldata newcid, string calldata oldcid, address sender, string calldata info) onlyOracleOwner external {
        if (statement == true) {
            address owner = dbOwner[oldcid];
            dbOwner[newcid] = owner;
            getcid[owner] = newcid;
            emit ReqState(reqID, sender, true, "Put data success.");
            emit NewCid(sender,newcid);
            reqStatement[reqID] = true;
        } else {
            emit ReqState(reqID, sender, false, info);
            reqStatement[reqID] = false;
        }
    }

    //get data from collection
    function Get(string calldata cid, string calldata colName, bytes calldata recordID, string calldata callBack) external override colIsExsist(cid, colName) allowQuery(cid) {
        emit get(CurrentReqID++, cid, colName, recordID, callBack, msg.sender);
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
    function Index(string calldata cid, string calldata colName, string calldata idx) external colIsExsist(cid, colName) allowWrite(cid) override {
        uint reqID = CurrentReqID;
        CurrentReqID++;
        emit index(reqID, cid, colName, idx, msg.sender);
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
    function Search(string calldata cid, string calldata colName, bytes calldata querys, string calldata callBack) external colIsExsist(cid, colName) allowQuery(cid) override {
        emit search(CurrentReqID++, cid, colName, querys, callBack, msg.sender);
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