#!/bin/bash
#
#  Created by Pequena Anta
#
#
# Usage ./issueTimeIntructionTransaction NumberOfTransactions 



round=$1
clients=$2

function getBlock() {
    docker exec cli "peer channel fetch 7 7block-${clients}clients-${round}round.block -c mychannel"
    docker exec cli "configtxlator proto_decode --input 7block-${clients}clients-${round}round.block --type common.Block --output 7block-${clients}clients-${round}round_read.block"
    docker exec cli "peer channel fetch newest newestblock-${clients}clients-${round}round.block -c mychannel"
    docker exec cli "configtxlator proto_decode --input newestblock-${clients}clients-${round}round.block --type common.Block --output newestblock-${clients}clients-${round}round_read.block"

    docker cp cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/newestblock-${i}clients-${index}round_read.block ~gta/gustavo/secTry/access-control-network/scripts/results/blocks &
    docker cp cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/7block-${i}clients-${index}round_read.block ~gta/gustavo/secTry/access-control-network/scripts/results/blocks &

}

getBlock
