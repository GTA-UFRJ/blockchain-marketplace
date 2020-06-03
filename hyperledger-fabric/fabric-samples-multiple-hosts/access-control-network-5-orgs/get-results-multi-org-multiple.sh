#!/bin/bash
# Created by Lucas Airam
#
#This scripts runs the scenario where multiple organization hosts multiples clients. The script issues 10000 advertisement transactions and times the emission to get the transaction rate at the client side

for z in 50 500 5000
do
for k in $(seq 1 10);
do
    for i in 1; 
    do
        printf "\n round $k start for $i clis start"
        #change the docker file to the corresonding client number in variable i
        ./credentials.sh >> /dev/null
        sleep 2
        ./activate_containers.sh $i >> /dev/null 
        sleep 2
        docker logs -f orderer2.example.com >> ./scripts/results/ordererlogs/''$i''clients-''$index''round.txt 2>&1 &
        #issues 10000 transactions
        ssh ipanema "docker exec cli1 \"./scripts/script.sh\"" >> /dev/null 
        #issues 10000 transactions in all running clients container
        printf "\nexecute transactions\n"
        then
            for counter in $(seq 0 $((i-1)));
            do
                ssh ipanema 'i='$i'; k='$k'; z='$z'; cmd="scripts/multiple-clients-per-org.sh $z $i $k"; docker exec cli$((1+5*i)) $cmd &' &
                ssh praiavermelha 'i='$i'; k='$k'; z='$z'; cmd="scripts/multiple-clients-per-org.sh $z $i $k"; docker exec cli$((2+5*i)) $cmd &' &
                ssh flamengo 'i='$i'; k='$k'; z='$z'; cmd="scripts/multiple-clients-per-org.sh $z $i $k"; docker exec cli$((3+5*i)) $cmd &' &
                ssh portuguesa 'i='$i'; k='$k'; z='$z'; cmd="scripts/multiple-clients-per-org.sh $z $i $k"; docker exec cli$((4+5*i)) $cmd &' &
                ssh leblon 'i='$i'; k='$k'; z='$z'; cmd="scripts/multiple-clients-per-org.sh $z $i $k"; docker exec cli$((5+5*i)) $cmd &' &
            done
        printf "\n sleeping again"
        sleep $(($i*$z/4))
        echo " end issue transactions\n" 
   #     ssh ipanema 'cd ~gta/gustavo/secTry/access-control-network-5-orgs; export index='$index'; cmd="scripts/getBlocks.sh $index $i"; docker exec cli $cmd &' &
    #    ssh ipanema "export index=$index; export i=$i; docker cp cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/newestblock-''$i''clients-''$index''round_read.block ~gta/gustavo/secTry/access-control-network/scripts/results/blocks"
     #   ssh ipanema "export index=$index; export i=$i; docker cp cli:/opt/gopath/src/github.com/hyperledger/fabric/peer/7block-'$i'clients-'$index'round_read.block ~gta/gustavo/secTry/access-control-network/scripts/results/blocks"
    printf "\n destroying network\n"
    ./desactivate_containers.sh $i   >> /dev/null
    printf "\nend round $k with $i clis"
    done;
done;
done;
