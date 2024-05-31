
# Prolly Tree Oracle

## What
* EVM oracle
* Based on [ipld prolly tree indexer](https://github.com/RangerMauve/ipld-prolly-indexer)
* Enable contracts to create/read/update/delete to prolly trees.
* Index over collections to enable faster queries
* charge gas or eth for transactions (e.g. if the field you query isn't indexed, charge more)
* maybe gas costs per lPLD block loaded? That way people deploying oracles can set theiown costs.
* allow querying by existing cid
* advertise prolly tree root on lPFS network to query offchain
* oracle publishes tree under ipns
* can we get read/write access gating based on wallets or multisigs?

## API:


* create(owner wallet, prolly config params)=>id

* allowWrite(wallet)

* put(id,collection,JSON document, primaryKey?) => primaryKey

* index(id, collection, fields)

* get(id, collection, primaryKey)->JSON

* search(id, collection, query {equals, sort, limit, skip})

## Architecture
Blockchain and decentralized applications (DApps) are increasingly important for establishing trust and transparency in data storage and computing. 
However, on-chain transactions are often costly and slow. To overcome this challenge, off-chain nodes can be used to store and compute data. 
Blockchain oracles are introduced and implemented in the form of application programming interfaces (APIs) to connect the real world to the blockchain. 
Prolly Tree Oracle connects the IPLD Prolly Tree storage service to the blockchain, allowing users to store large amounts of data off-chain at a fraction of the cost while enjoying the transparency and security brought by the blockchain.

![](/sources/Architecture.png)


## API Implementation:
* **function Create(string calldata dbName, string calldata colName, string calldata primaryKey)**  
Creates a new collection named `colName` in the database pointed to by the provided `dbName`.
If two users register the same `dbName` successively, the owner of this database is the first user. 
When the second user registers, it is necessary to determine whether it has the permission to the database.


* **function AllowWrite(address to)**  
Grant user `to` read and write permissions on your database.


* **function Put(string calldata dbName, string calldata colName, bytes calldata data)**  
  Save the data `data` to the `colName` collection in the database pointed to by `dbName`.
  

* **function Get(string calldata dbName, string calldata colName, bytes calldata recordID, string calldata callBack)**  
  Get data from the `colName` collection in the database pointed to by `dbName` according to `recordID`. You need to provide a callback function to receive the data. 
  For example, the callback function I defined in the consumer contract is:`function CallBackFunc(bytes calldata data)`.  So, the callBack parameter of string type is ``CallBackFunc(bytes)``.  


* **function Index(string calldata dbName, string calldata colName, string calldata idx)**  
  Add index `idx` to the `colName` collection in the database pointed to by `dbName`


* **function Search(string calldata dbName, string calldata colName, bytes calldata querys, string calldata callBack)**  
  Use `index.Query` to query data from the `colName` collection in the database pointed to by `dbName`. When using it, you can build one or more `index.Query` objects, and serialize the `index.Query` object array into `json` as the parameter `queries`. Query methods:`{equal, compare, sort, limit, skip}` 

  Similarly, a callback function needs to be provided to receive data  


* **function GetRootCid(string calldata dbName, string calldata callBack)**

  Get the `RootCid` of the database of the method caller

  Similarly, a callback function needs to be provided to receive an array of collection indexes


* **function GetIndex(string calldata dbName, string calldata colName, string calldata callBack)**

  Get the index of the `colName` collection in the database pointed to by `dbName`

  Similarly, a callback function needs to be provided to receive an array of collection indexes


* **function GetCol(string calldata dbName, string calldata callBack)**

  Get all collections in the database pointed to by `dbName`
  Similarly, a callback function needs to be provided to receive an array of collection names


# Usage
### [Click here to learn how to start the service](https://www.youtube.com/watch?v=_T4nLVJaQ5k)
## Consumer Contract

The **Consumer Contract** used for testing is located at `contracts/test.sol`

Compile the **Consumer Contract** to obtain the ABI file, and use the ***[abigen](https://geth.ethereum.org/docs/tools/abigen)*** to generate **Golang package** corresponding to the contract:
`abigen --abi={ Path to abi.json } --pkg={ Package name } --out={ The path to the generated .go file }`

## Golang Client
Use a Golang test program `golangServer/test/Oracle_test.go` to call the request function in the Consumer Contract for testing the oracle.
