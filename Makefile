gen:
	cd ./contracts && solcjs --abi oracle.sol -o build 
	cp ./contracts/build/oracle_sol_Oracle.abi ./golangServer/Oracle/abi.json
	abigen --abi=./golangServer/Oracle/abi.json --pkg=Oracle --out=./golangServer/Oracle/Oracle.go
