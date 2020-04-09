#!/bin/bash
# Created by Lucas Airam, based on get-results-per-org.sh by Gustavo Camilo
#
#This scripts runs the scenario where one organization hosts multiples clients. The script issues 10000 advertisement transactions and times the emission to get the transaction rate at the client side
#c=6
#d=21
#e=79
sed -i '477s/.*/COMPOSE_FILE=docker-compose-16cli-perorg.yaml/' byfn.sh        
. byfn.sh up #>> /dev/null 2>&1
printf "\n sleeping"
sleep 60
for z in 50
do
for i in 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30;
do 
    for k in 1; #$(seq 3 10);
    do
        printf "\n round $k start for $i clis start"
        #change the docker file to the corresonding client number in variable i
        cmd="scripts/multiple-clients-per-org.sh $z $i $k"
        printf "\n execute transactions\n"
        docker exec cli $cmd &
        #issues 5000 transactions in all running clients container
        for counter in $(seq 2 $i);
            do
                docker exec cli$counter $cmd &
            done
        #waits for the transactions to finish to end the docker network
        printf "\n sleeping again"
        sleep $(($i*$z/13))
        #sleeps longer for more than 4 clients so everyone can send the transactions
#        if [ $(($i*5)) -gt $c ]
#        then
#            sleep 50
#            if [ $(($i*5)) -gt $d ]
#            then
#                sleep 60
#                if [ $(($i*5)) -gt $e ]
#                then
#                    sleep 60
#                fi
#            fi
#        fi
        printf "\n destroying network\n"
#        docker rm -f $(docker ps -aq) >> /dev/null 2>&1 && docker rmi -f $(docker images | grep dev | awk '{print $3}') >> /dev/null 2>&1 && docker volume prune -f >> /dev/null 2>&1
#        . byfn.sh down >> /dev/null 2>&1
        printf "\n end round $k with $i clis"
    done;
done;
done;
