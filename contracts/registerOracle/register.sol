pragma solidity ^0.8.0;

contract Register {
    mapping (string => string) private oracleList;
    string[] private keys;

    address private owner;

    constructor(){
        owner = msg.sender;
    }

    function set(string calldata key, string calldata value) public {
        require(msg.sender == owner,"Only the owner can call this function. ");
        require(bytes(oracleList[key]).length == 0,"This key has been registered");
        keys.push(key);
        oracleList[key] = value;
    }

    function remove(string calldata key) public {
        require(msg.sender == owner,"Only the owner can call this function. ");
        delete oracleList[key];
        for (uint i = 0; i < keys.length; i++) {
            if (keccak256(bytes(keys[i])) == keccak256(bytes(key))) {
                keys[i] = keys[keys.length - 1];
                keys.pop();
                break;
            }
        }
    }

    function getKeys() public view returns (string[] memory) {
        return keys;
    }

    function getValue(string calldata key) public view returns (string memory) {
        return oracleList[key];
    }
}