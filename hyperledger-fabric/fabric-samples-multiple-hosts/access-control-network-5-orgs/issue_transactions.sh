#!/bin/bash
z=$1
i=$2
k=$3
cmd="scripts/multiple-clients-per-org.sh $z $i $k"
        for counter in $(seq 0 15);
        do
            docker exec cli$((1+5*$counter)) $cmd &
        done
