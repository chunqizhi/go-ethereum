#!/bin/sh
geth --datadir /data init /data/genesis.json

geth --datadir /data --networkid 4857 --allow-insecure-unlock --ws --wsaddr 0.0.0.0 --wsorigins=* --rpc --rpcaddr 0.0.0.0 --rpcapi admin,eth,miner,web3,personal,net,txpool --rpcvhosts=*

