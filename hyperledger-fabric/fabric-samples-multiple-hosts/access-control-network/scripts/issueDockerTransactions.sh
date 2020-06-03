#!/bin/bash
#Created by Gustavo
#A script that issues transactions by a remote host

clients=$1
transactions=$2

cmd="./scripts/multiple-clients-same-org.sh $transactions $clients" 

docker exec cli $cmd &
for counter in $(seq 2 $clients);
do
    docker exec cli$counter $cmd &
    now=$(date)
    echo "Cliente $counter emitiu transacoes as $now" >> ~gta/gustavo/secTry/access-control-network/scripts/results/prints/arquivo.txt
done
