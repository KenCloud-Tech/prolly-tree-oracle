// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Oracle} from "./oracle.sol";

contract OracleTest{
    Oracle public or;
    function setOracle(address add) public{
        or = Oracle(add);
    }
    //allow a address write
    function AllowWrite(string calldata dbName, address to) external{
        or.AllowWrite(dbName,to);
    }

    //creat a new db
    function Create(string calldata dbName, string calldata primaryKey) external{
        or.Create(dbName,primaryKey);
    }

    // put data to db
    function Put(string calldata dbName, bytes calldata data) external{
        or.Put(dbName,data);
    }

    //get data from db
    function Get(string calldata dbName, bytes calldata recordID, string calldata callBack) external{
        or.Get(dbName,recordID,callBack);
    }

    //creat index
    function Index(string calldata dbName, string calldata Key) external{
        or.Index(dbName,Key);
    }
    
    //query by {equals, compare, sort, limit, skip}
    function Search(string calldata dbName, Oracle.SearchController calldata Val, string calldata Method, string calldata callBack) external{
        or.Search(dbName,Val,Method,callBack);
    }

    // catch data and emit event
    event sendData(bytes data);
    function catchData(bytes calldata data) external{
        emit sendData(data);
    }

}