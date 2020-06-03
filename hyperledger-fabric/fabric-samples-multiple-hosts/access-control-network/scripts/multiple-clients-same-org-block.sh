#!/bin/bash
#
#  Created by Pequena Anta
#
#
# Usage ./issueTimeIntructionTransaction NumberOfTransactions 


export CHANNEL_NAME=mychannel

transaction=$1
clients=$2
blockTransaction=$3

function issueTransaction() {
    { time for transactionNumber in `seq 1 $transaction`;
    do
        peer chaincode invoke -n mycc -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -c '{"Args":["issueAdvertisement","Dados IoT","Dados de Sensores IoT","10","IoT","10.0.0.1","Org1"]}' > /dev/null 2>&1


    done ; } 2>> /opt/gopath/src/github.com/hyperledger/fabric/peer/scripts/results/multiple-clients-same-org/"$clients"clients-same-org-"$blockTransaction"block.txt  
}

issueTransaction 
printf "\n\n"
