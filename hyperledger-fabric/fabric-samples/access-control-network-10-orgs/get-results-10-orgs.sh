#!/bin/bash
# Created by Gustavo Camilo
#
#This scripts runs the scenario where one organization hosts multiples clients. The script issues 10000 advertisement transactions and times the emission to get the transaction rate at the client side
b=32

for index in $(seq 1 10)
do
    #. byfn.sh up >> /dev/null 2>&1
    . byfn.sh up
    #issues 10000 transactions
    cmd="scripts/1-client-10-orgs.sh 5000 $i"
    docker exec cli $cmd &

    #waits for the transactions to finish to end the docker network
    sleep 3600
    . byfn.sh down >> /dev/null 2>&1
done;
