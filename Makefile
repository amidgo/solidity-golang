build-contract:
	solc --abi ./contracts/Database.sol > ./contracts/Database.abi
	solc --bin ./contracts/Database.sol > ./contracts/Database.bin
go-contract:
	abigen --bin=./contracts/Database.bin --abi=./contracts/Database.abi --pkg=contracts --out=./contracts/Database.go
node1:
	geth --datadir eth-net/node1 --networkid 12345 --port 30306 --bootnodes enode://742c4afbb589817939dc8e6382e60026c0a49224111f446da1548a473aaf749a067e8c03d3c6f30ceb1b0467f59eff69bae076d307c7072e8abfc3ab1354ebbe@127.0.0.1:0?discport=30305 --http --http.port 1111 --http.api net,web3,eth,debug,miner,personal,admin,txpool --allow-insecure-unlock --unlock 0x87B601B72D92D5D4064888D51cF84f40b12B9D27 --password eth-net/coinbase.txt
node2:
	geth --datadir eth-net/node2 --networkid 12345 --port 30307 --bootnodes enode://742c4afbb589817939dc8e6382e60026c0a49224111f446da1548a473aaf749a067e8c03d3c6f30ceb1b0467f59eff69bae076d307c7072e8abfc3ab1354ebbe@127.0.0.1:0?discport=30305 --unlock 0xf0Bd05F21e09439C4CCE66D8aA8f4bcAB93F4F7f --password eth-net/miner.txt --mine --miner.etherbase 0xf0Bd05F21e09439C4CCE66D8aA8f4bcAB93F4F7f --authrpc.port 8550 --http -http.port 2222 --http.api net,web3,eth,debug,miner,personal,admin,txpool --allow-insecure-unlock
bootnode:
	bootnode -nodekey eth-net/boot.key -addr :30305
refresh-nodes:
	rm -r eth-net/node1/*
	rm -r eth-net/node2/*
	rm -r eth-net/testkeystore/*
	geth init --datadir eth-net/node1 eth-net/genesis.json
	geth init --datadir eth-net/node2 eth-net/genesis.json
	cp -r eth-net/keystore/. eth-net/node1/keystore
	cp -r eth-net/keystore/. eth-net/node2/keystore
	cp -r eth-net/keystore/. eth-net/testkeystore
test-node:
	geth --http  --http.api "web3,eth,debug,net,personal" --allow-insecure-unlock --dev console --keystore ./eth-net/testkeystore --unlock 0x87B601B72D92D5D4064888D51cF84f40b12B9D27 --password eth-net/coinbase.txt

