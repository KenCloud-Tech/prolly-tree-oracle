export const oracleAbi=[
    {
        "inputs": [
            {
                "internalType": "address",
                "name": "to",
                "type": "address"
            }
        ],
        "name": "AllowWrite",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "statement",
                "type": "bool"
            },
            {
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "info",
                "type": "string"
            }
        ],
        "name": "CreatRsp",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "basefee",
                "type": "uint256"
            },
            {
                "internalType": "uint256",
                "name": "bytefee",
                "type": "uint256"
            }
        ],
        "stateMutability": "payable",
        "type": "constructor"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            }
        ],
        "name": "CatchData",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "user",
                "type": "address"
            },
            {
                "indexed": false,
                "internalType": "bool",
                "name": "statement",
                "type": "bool"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "info",
                "type": "string"
            }
        ],
        "name": "ReqState",
        "type": "event"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "primaryKey",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "owner",
                "type": "address"
            }
        ],
        "name": "create",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "primaryKey",
                "type": "string"
            }
        ],
        "name": "Create",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "bytes",
                "name": "recordID",
                "type": "bytes"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "get",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "internalType": "bytes",
                "name": "recordID",
                "type": "bytes"
            },
            {
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            }
        ],
        "name": "Get",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "getCol",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            }
        ],
        "name": "GetCol",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "ReqID",
                "type": "uint256"
            }
        ],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "statement",
                "type": "bool"
            },
            {
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            },
            {
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            },
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "info",
                "type": "string"
            }
        ],
        "name": "GetColRsp",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "getIndex",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            }
        ],
        "name": "GetIndex",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "ReqID",
                "type": "uint256"
            }
        ],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "statement",
                "type": "bool"
            },
            {
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            },
            {
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            },
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "info",
                "type": "string"
            }
        ],
        "name": "GetIndexRsp",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "getRootCid",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            }
        ],
        "name": "GetRootCid",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "ReqID",
                "type": "uint256"
            }
        ],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "statement",
                "type": "bool"
            },
            {
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            },
            {
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            },
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "info",
                "type": "string"
            }
        ],
        "name": "GetRootCidRsp",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "statement",
                "type": "bool"
            },
            {
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            },
            {
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            },
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "info",
                "type": "string"
            }
        ],
        "name": "GetRsp",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "url",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "format",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "import_from_url",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "url",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "format",
                "type": "string"
            }
        ],
        "name": "Import_from_url",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "statement",
                "type": "bool"
            },
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "info",
                "type": "string"
            }
        ],
        "name": "Import_from_url_Rsp",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "Key",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "index",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "idx",
                "type": "string"
            }
        ],
        "name": "Index",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "statement",
                "type": "bool"
            },
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "info",
                "type": "string"
            }
        ],
        "name": "IndexRsp",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "put",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            }
        ],
        "name": "Put",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "statement",
                "type": "bool"
            },
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "info",
                "type": "string"
            }
        ],
        "name": "PutRsp",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "anonymous": false,
        "inputs": [
            {
                "indexed": false,
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "bytes",
                "name": "querys",
                "type": "bytes"
            },
            {
                "indexed": false,
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            },
            {
                "indexed": false,
                "internalType": "address",
                "name": "sender",
                "type": "address"
            }
        ],
        "name": "search",
        "type": "event"
    },
    {
        "inputs": [
            {
                "internalType": "string",
                "name": "dbName",
                "type": "string"
            },
            {
                "internalType": "string",
                "name": "colName",
                "type": "string"
            },
            {
                "internalType": "bytes",
                "name": "querys",
                "type": "bytes"
            },
            {
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            }
        ],
        "name": "Search",
        "outputs": [],
        "stateMutability": "payable",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "reqID",
                "type": "uint256"
            },
            {
                "internalType": "bool",
                "name": "statement",
                "type": "bool"
            },
            {
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            },
            {
                "internalType": "string",
                "name": "callBack",
                "type": "string"
            },
            {
                "internalType": "address",
                "name": "sender",
                "type": "address"
            },
            {
                "internalType": "string",
                "name": "info",
                "type": "string"
            }
        ],
        "name": "SearchRsp",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "withdraws",
        "outputs": [],
        "stateMutability": "nonpayable",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "baseGasFee",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [],
        "name": "gasPerByte",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "bytes",
                "name": "data",
                "type": "bytes"
            }
        ],
        "name": "getPrice",
        "outputs": [
            {
                "internalType": "uint256",
                "name": "",
                "type": "uint256"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    },
    {
        "inputs": [
            {
                "internalType": "uint256",
                "name": "ReqID",
                "type": "uint256"
            }
        ],
        "name": "GetReqState",
        "outputs": [
            {
                "internalType": "bool",
                "name": "",
                "type": "bool"
            }
        ],
        "stateMutability": "view",
        "type": "function"
    }
]