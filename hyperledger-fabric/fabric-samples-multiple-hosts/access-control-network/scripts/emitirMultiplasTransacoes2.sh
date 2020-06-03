#!/bin/bash
#
#  Created by Pequena Anta
#
#
# Usage ./issueTimeIntructionTransaction NumberOfTransactions 


export CHANNEL_NAME=mychannel

transaction=$1
function issueTransaction() {
    start_time=$SECONDS
    start_time=$(($(date +%s%N)/1000000))
    time for transactionNumber in `seq 1 $transaction`;
    do
        peer chaincode invoke -n mycc -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -c '{"Args":["issueAdvertisement","Dados IoT","Dados de Sensores IoT","10","IoT","10.0.0.1","Org1"]}' > /dev/null 2>&1


    done
    end_time=$SECONDS
   end_time=$(($(date +%s%N)/1000000))

    printf "\n\n"
    echo ------------------------ RESULTS -------------------------------
    printf "\n\n"
    
    echo $transactionNumber response transactions issued
    echo Duration of simulation: $(($end_time - $start_time)) miliseconds
    
    printf "\n\n"
    echo -----------------------------------------------------------------
    printf "\n\n"
}

issueTransaction 
printf "\n\n"
