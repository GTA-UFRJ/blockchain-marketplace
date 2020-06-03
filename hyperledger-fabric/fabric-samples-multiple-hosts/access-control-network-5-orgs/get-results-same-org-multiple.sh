#!/bin/bash
# Created by Gustavo Camilo
#
#This scripts runs the scenario where one organization hosts multiples clients. The script issues 10000 advertisement transactions and times the emission to get the transaction rate at the client side
a=1
b=2
c=4
d=8
e=16
f=32


for index in $(seq 1 1)
do
    for i in 4; 
    do
        #change the docker file to the corresonding client number in variable i
        . byfn.sh generate >> /dev/null 2>&1
        scp -r crypto-config gta@ipanema:~gta/gustavo/secTry/access-control-network-5-orgs >> /dev/null 2>&1
        scp -r channel-artifacts gta@ipanema:~gta/gustavo/secTry/access-control-network-5-orgs >> /dev/null 2>&1
        scp -r crypto-config gta@praiavermelha:~gta/gustavo/secTry/access-control-network-5-orgs >> /dev/null 2>&1
        scp -r channel-artifacts gta@praiavermelha:~gta/gustavo/secTry/access-control-network-5-orgs >> /dev/null 2>&1
        scp -r crypto-config gta@flamengo:~gta/gustavo/secTry/access-control-network-5-orgs >> /dev/null 2>&1
        scp -r channel-artifacts gta@flamengo:~gta/gustavo/secTry/access-control-network-5-orgs >> /dev/null 2>&1
        scp -r crypto-config gta@portuguesa:~gta/gustavo/secTry/access-control-network-5-orgs >> /dev/null 2>&1
        scp -r channel-artifacts gta@portuguesa:~gta/gustavo/secTry/access-control-network-5-orgs >> /dev/null 2>&1
        scp -r crypto-config gta@leblon:~gta/gustavo/secTry/access-control-network-5-orgs >> /dev/null 2>&1
        scp -r channel-artifacts gta@leblon:~gta/gustavo/secTry/access-control-network-5-orgs >> /dev/null 2>&1
        docker-compose -f docker-compose-etcdraft2-modified.yaml up -d >> /dev/null 2>&1
        sleep 3
        dockrer logs -f orderer2.example.com >> ./scripts/results/ordererlogs/''$i''clients-''$index''round.txt 2>&1 &
        ssh praiavermelha "cd ~gta/gustavo/secTry/access-control-network; docker-compose -f base/docker-compose-org2.yaml up -d >> /dev/null 2>&1"
        sleep 2
        ssh ipanema "cd ~gta/gustavo/secTry/access-control-network; export i=$i; docker-compose -f docker-compose-''$i''cli-sameorg-modified.yaml up -d >> /dev/null 2>&1"
        sleep 2
        #issues 10000 transactions
        ssh ipanema "docker exec cli \"./scripts/script.sh\" &"
        sleep 120
        ssh ipanema 'i='$i'; cmd="scripts/multiple-clients-same-org.sh 500 $i"; docker exec cli $cmd &' &
        #issues 10000 transactions in all running clients container
        if [ $i -ge $b ]
        then
            for counter in $(seq 2 $i);
            do
                ssh ipanema 'i='$i'; counter='$counter'; cmd="scripts/multiple-clients-same-org.sh 500 $i"; docker exec cli$counter $cmd &' &
            done
        fi
        #waits for the transactions to finish to end the docker network
        sleep 30
        #sleeps longer for more than 4 clients so everyone can send the transactions
        if [ $i -gt $a ]
        then
            sleep 30
            if [ $i -gt $b ]
            then
                sleep 30
                if [ $i -gt $c ]
                then
                    sleep 30
                    if [ $i -gt $d ]
                    then
                        sleep 30
                        if [ $i -gt $e ]
                        then
                            sleep 30
                            sleep 30
                        fi
                    fi
                fi
            fi
        fi
        ssh ipanema 'cd ~gta/gustavo/secTry/access-control-network; export index='$index'; cmd="scripts/getBlocks.sh $index $i"; docker exec cli $cmd &' &
        ssh ipanema "export index=$index; export i=$i; docker cp cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/newestblock-''$i''clients-''$index''round_read.block ~gta/gustavo/secTry/access-control-network/scripts/results/blocks"
        ssh ipanema "export index=$index; export i=$i; docker cp cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/7block-'$i'clients-'$index'round_read.block ~gta/gustavo/secTry/access-control-network/scripts/results/blocks"
        . byfn.sh down >> /dev/null 2>&1
        ssh praiavermelha "cd ~gta/gustavo/secTry/access-control-network; . byfn.sh down >> /dev/null 2>&1"
        ssh ipanema "cd ~gta/gustavo/secTry/access-control-network; . byfn.sh down >> /dev/null 2>&1"
    done;
done;
