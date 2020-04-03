#!/bin/bash
# Created by Gustavo Camilo
#
#This scripts runs the scenario where one organization hosts multiples clients. The script issues 10000 advertisement transactions and times the emission to get the transaction rate at the client side
b=32

for index in $(seq 1 10)
do
    for i in 32 64; 
    do
        #change the docker file to the corresonding client number in variable i
        sed -i '430s/.*/COMPOSE_FILE=docker-compose-'"$i"'cli-sameorg.yaml/' byfn.sh        
        . byfn.sh up >> /dev/null 2>&1
        #issues 10000 transactions
        cmd="scripts/multiple-clients-same-org.sh 5000 $i"
        docker exec cli $cmd &
        #issues 10000 transactions in all running clients container
        if [ $i -ge $b ]
        then
            for counter in $(seq 2 $i);
            do
                docker exec cli$counter $cmd &
            done
        fi
        #waits for the transactions to finish to end the docker network
        sleep 3600
        #sleeps longer for more than 4 clients so everyone can send the transactions
        if [ $i -gt $b ]
        then
            sleep 3600
        fi
        . byfn.sh down >> /dev/null 2>&1
    done;
done;
