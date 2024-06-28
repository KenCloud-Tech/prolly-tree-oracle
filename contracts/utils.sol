// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IOracle} from "./I_oracle.sol";

contract util is IOracle {
    // A string is empty?
    function isEmptyString(string memory str) internal pure returns (bool) {
        return bytes(str).length == 0;
    }
    // Estimated price
    function getPrice(bytes calldata data) public view returns (uint) {
        return data.length * gasPerByte;
    }
    // Pay for service
    function pay(uint fee) internal{
        fee = fee + baseGasFee;
        require(msg.value >= fee, "Insufficient gas fee");
        payable(msg.sender).transfer(msg.value - fee);
    }
    // Oracle owner withdraws gas
    function withdraws() external onlyOracleOwner{
        require(address(this).balance > 0, "Balance = 0 !");
        owner.transfer(address(this).balance);
    }




    //get collections
    event getCol(uint reqID, string dbName, string callBack, address sender);
    function GetCol(string calldata dbName, string calldata callBack) external allowQuery(dbName) payable returns(uint ReqID){
        pay(0);
        emit getCol(CurrentReqID++, dbName, callBack, msg.sender);
        return CurrentReqID;
    }
    function GetColRsp(uint reqID, bool statement, bytes calldata data, string calldata callBack, address sender, string calldata info) onlyOracleOwner external {
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

    //get indexes
    event getIndex(uint reqID, string dbName, string colName, string callBack, address sender);
    function GetIndex(string calldata dbName, string calldata colName, string calldata callBack) external allowQuery(dbName) payable returns(uint ReqID){
        pay(0);
        require(cols[dbName][colName] != false, "This collection has not been created yet");
        emit getIndex(CurrentReqID++, dbName, colName, callBack, msg.sender);
        return CurrentReqID;
    }

    function GetIndexRsp(uint reqID, bool statement, bytes calldata data, string calldata callBack, address sender, string calldata info) onlyOracleOwner external {
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

    //get RootCid
    event getRootCid(uint reqID, string dbName, string callBack, address sender);
    function GetRootCid(string calldata dbName, string calldata callBack) external payable returns(uint ReqID){
        pay(0);
        require(dbOwner[dbName]!=address(0),"This db is not exist.");
        emit getRootCid(CurrentReqID++, dbName, callBack, msg.sender);
        return CurrentReqID;
    }

    function GetRootCidRsp(uint reqID, bool statement, bytes calldata data, string calldata callBack, address sender, string calldata info) onlyOracleOwner external {
        if (statement == true) {
            (bool OK,) = sender.call(abi.encodeWithSignature(callBack, data));
            if (OK) {
                emit ReqState(reqID, sender, true, "Get RootCid success.");
                reqStatement[reqID] = true;
                return;
            }
        }

        emit ReqState(reqID, sender, false, info);
        reqStatement[reqID] = false;
    }

}