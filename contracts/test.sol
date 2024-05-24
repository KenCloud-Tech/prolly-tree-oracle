// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Oracle} from "./oracle.sol";

contract OracleTest{
    Oracle public or;
    function setOracle(address add) public{
        or = Oracle(add);
    }
    //allow a address write
    function AllowWrite(address to) external{
        or.AllowWrite(to);
    }

    //creat a new db
    function Create(string calldata cid, string calldata colName, string calldata primaryKey) external{
        or.Create(cid, colName,primaryKey);
    }

    // put data to db
    function Put(string calldata cid, string calldata colName, bytes calldata data) external{
        or.Put(cid, colName,data);
    }

    //get data from db
    function Get(string calldata cid, string calldata colName, bytes calldata recordID, string calldata callBack) external{
        or.Get(cid, colName,recordID,callBack);
    }

    //creat index
    function Index(string calldata cid, string calldata colName, string calldata Key) external{
        or.Index(cid, colName,Key);
    }
    
    //query by {equals, compare, sort, limit, skip}
    function Search(string calldata cid, string calldata colName,bytes calldata querys, string calldata callBack) external{
        or.Search(cid, colName,querys,callBack);
    }

    function GetCol(string calldata cid, string calldata callBack) external {
        or.GetCol(cid,callBack);
    }

    function GetIndex(string calldata cid, string calldata colName, string calldata callBack) external {
        or.GetIndex(cid,colName,callBack);
    }

    function getRootCid() view public returns(string memory){
        return or.getRootCid();
    }

    // catch data and emit event
    event CatchData(bytes data);
    function CBFunc(bytes calldata data) external{
        //You can process the obtained data here.
        emit CatchData(data);
    }

}