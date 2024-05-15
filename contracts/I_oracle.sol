// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Permission} from "./interfaces/permission.sol";

contract IOracle{

    // struct ReqController{
    //     uint reqID;
    //     address userAdd;//requester`s address
    //     string callBack;//callback function
    //     bool statement;// true or false
    // }
    bytes32 equal = keccak256(abi.encodePacked("equal"));
    bytes32 compare = keccak256(abi.encodePacked("compare"));
    bytes32 sort = keccak256(abi.encodePacked("sort"));
    bytes32 limit = keccak256(abi.encodePacked("limit"));
    bytes32 skip = keccak256(abi.encodePacked("skip"));

    uint internal CurrentReqID;
    event ReqState(uint reqID, address user, bool statement, string info);

    //contract owner
    address private owner;
    // A mapping of db ids to db owner`s addresses.
    mapping(string=>address) internal dbOwner;
    // A mapping of user address to db control permission.
    mapping(address=>mapping(string=>Permission)) internal permission;
    // A mapping of user address to request statement
    // mapping(address => mapping(uint=>ReqController)) internal reqStatement;

    constructor() {
            owner = msg.sender;
            CurrentReqID = 0;
        }
    
    modifier onlyOracleOwner() {
        require(msg.sender == owner, "Only the owner can call this function");
        _;
    }
    modifier onlyDbOwner(string calldata dbName) {
        require(msg.sender == dbOwner[dbName], "Only the db owner can call this function");
        _;
    }
    // modifier allowInsert(string calldata dbName){
    //     require(permission[msg.sender][dbName].allowInsert == true,"You do not have permission to insert");
    //     _;
    // }
    // modifier allowUpdate(string calldata dbName){
    //     require(permission[msg.sender][dbName].allowUpdate == true,"You do not have permission to update");
    //     _;
    // }
    modifier allowWrite(string calldata dbName){
        require(permission[msg.sender][dbName].allowWrite == true,"You do not have permission to write");
        _;
    }
    modifier allowQuery(string calldata dbName){
        require(permission[msg.sender][dbName].allowQuery == true,"You do not have permission to query");
        _;
    }
    modifier allowDelete(string calldata dbName){
        require(permission[msg.sender][dbName].allowDelete == true,"You do not have permission to delete");
        _;
    }
}