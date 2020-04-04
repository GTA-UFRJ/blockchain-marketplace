#!/bin/bash
# Created by Lucas Airam, based on get-results-per-org.sh by Gustavo Camilo
#
#This scripts runs the scenario where one organization hosts multiples clients. The script issues 10000 advertisement transactions and times the emission to get the transaction rate at the client side
b=2
c=4
d=8
e=16

for index in $(seq 1 10)
do
    for i in 1 2 4 8 16; 
    do
        echo ENTREI AQUI $i
        #change the docker file to the corresonding client number in variable i
        sed -i '430s/.*/COMPOSE_FILE=docker-compose-'"$i"'cli-perorg.yaml/' byfn.sh        
        . byfn.sh up #>> /dev/null 2>&1
        #issues 5000 transactions
        cmd="scripts/multiple-clients-per-org.sh 5000 $i"
        docker exec cli $cmd &
        #issues 5000 transactions in all running clients container
        if [ $i -ge $b ]
        then
            for counter in $(seq 2 $(($i*5)));
            do
                docker exec cli$counter $cmd &
            done
        fi
        #waits for the transactions to finish to end the docker network
        sleep 500
        #sleeps longer for more than 4 clients so everyone can send the transactions
        if [ $i -gt $c ]
        then
            sleep 500
            if [ $i -gt $d ]
            then
                sleep 500
                if [ $i -gt $e ]
                then
                    sleep 1000
                fi
            fi
        fi
        . byfn.sh down >> /dev/null 2>&1
    done;
done;
