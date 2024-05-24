// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Permission} from "./interfaces/permission.sol";

contract IOracle {


    uint internal CurrentReqID;

    event ReqState(uint reqID, address user, bool statement, string info);
    event NewCid(address owner,string cid);

    //contract owner
    address private owner;
    // A mapping of cid to db owner`s addresses. cid=>dbOwner address
    mapping(string => address) internal dbOwner;
    // A mapping of db owner`s addresses to cid. dbOwner address=>cid
    mapping(address=>string) getcid;
    // Mapping from cid to whether collection exists   dbOwner`s address=>collection=>isExist
    mapping(address => mapping(string => bool)) cols;
    // A mapping of user address to db control permission.  address=>dbOwner`s address=>permission
    mapping(address => mapping(address => Permission)) internal permission;
    // A mapping of user address to request statement
    mapping(uint => bool) internal reqStatement;

    function GetReqState(uint ReqID) public view returns (bool){
        return reqStatement[ReqID];
    }

    constructor() {
        owner = msg.sender;
        CurrentReqID = 1;
    }

    modifier onlyOracleOwner() {
        require(msg.sender == owner, "Only the oracle owner can call this function");
        _;
    }
    modifier onlyDbOwner() {
        require(msg.sender == dbOwner[getcid[msg.sender]], "Only the db owner can call this function");
        _;
    }
    modifier dbIsExsist(string calldata cid) {
        require(dbOwner[cid] != address(0), "Db is not exist");
        _;
    }
    modifier colIsExsist(string calldata cid, string calldata colName) {
        require(cols[dbOwner[cid]][colName] != false, "This collection has not been created yet");
        _;
    }
    // modifier allowInsert(string calldata cid){
    //     require(permission[msg.sender][cid].allowInsert == true,"You do not have permission to insert");
    //     _;
    // }
    // modifier allowUpdate(string calldata cid){
    //     require(permission[msg.sender][cid].allowUpdate == true,"You do not have permission to update");
    //     _;
    // }
    modifier allowWrite(string calldata cid){
        require(bytes(cid).length == 0 || permission[msg.sender][dbOwner[cid]].allowWrite == true, "You do not have permission to write");
        _;
    }
    modifier allowQuery(string calldata cid){
        require(permission[msg.sender][dbOwner[cid]].allowQuery == true, "You do not have permission to query");
        _;
    }
    modifier allowDelete(string calldata cid){
        require(permission[msg.sender][dbOwner[cid]].allowDelete == true, "You do not have permission to delete");
        _;
    }

}