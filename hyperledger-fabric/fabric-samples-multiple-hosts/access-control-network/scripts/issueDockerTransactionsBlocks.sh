#!/bin/bash
#Created by Gustavo
#A script that issues transactions by a remote host

clients=$1
transactions=$2
blockTransaction=$3

cmd="./scripts/multiple-clients-same-org-block.sh $transactions $clients $blockTransaction" 

docker exec cli $cmd &
for counter in $(seq 2 $clients);
do
    docker exec cli$counter $cmd &
done
