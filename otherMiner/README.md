# chaojigongshi-other-miner
cd ..

make geth

cd otherMiner

rm -rf data/*

../build/bin/geth --datadir ./data init genesis.json

nohup ../build/bin/geth --datadir ./data --networkid 4857 --allow-insecure-unlock --ws --wsaddr 0.0.0.0 --wsorigins=* --rpc --rpcaddr 0.0.0.0 --rpcapi admin,eth,miner,web3,personal,net,txpool --rpcvhosts=* > geth.log 2>&1 &

../build/bin/geth attach data/geth.ipc
