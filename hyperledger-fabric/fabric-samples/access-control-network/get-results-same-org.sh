#!/bin/bash
# Created by Gustavo Camilo
#
#This scripts runs the scenario where one organization hosts multiples clients. The script issues 10000 advertisement transactions and times the emission to get the transaction rate at the client side

for i in 1; 
do
    sed -i '430s/.*/COMPOSE_FILE=docker-compose-'"$i"'cli-sameorg.yaml/' byfn.sh        
    . byfn.sh up
    cmd="scripts/multiple-clients-same-org.sh 10000 $i"
    docker exec cli $cmd &
    echo "Sleeping for 60s..."
    sleep 60
    . byfn.sh down
done;
