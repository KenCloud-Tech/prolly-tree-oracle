// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IOracle} from "./I_oracle.sol";

contract util is IOracle {
    // A string is empty?
    function isEmptyString(string memory str) internal pure returns (bool) {
        return bytes(str).length == 0;
    }

    //get collections
    event getCol(uint reqID, string cid, string callBack, address sender);
    function GetCol(string calldata cid, string calldata callBack) external allowQuery(cid) {
        emit getCol(CurrentReqID++, cid, callBack, msg.sender);
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
    event getIndex(uint reqID, string cid, string colName, string callBack, address sender);
    function GetIndex(string calldata cid, string calldata colName, string calldata callBack) external colIsExsist(cid, colName) allowQuery(cid) {
        emit getIndex(CurrentReqID++, cid, colName, callBack, msg.sender);
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

    // get rootcid
    function GetRootCid() view public returns(string memory){
        return getcid[msg.sender];
    }
}