### remix连接本地文件夹<br>
`remixd -s ./contracts -u https://remix.ethereum.org:8080`

### 生成合约的go包<br>
`abigen --abi=golangServer/Oracle/abi.json --pkg=Oracle --out=golangServer/Oracle/Oracle.go`