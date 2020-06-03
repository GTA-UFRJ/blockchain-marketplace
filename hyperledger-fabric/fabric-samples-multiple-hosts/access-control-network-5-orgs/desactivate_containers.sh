#!/bin/bash

i=$1        

ssh ipanema "cd ~gta/gustavo/secTry/access-control-network-5-orgs; export i=$i; docker-compose -f docker-compose-''$i''cli-multiorg1.yaml down; rm docker-compose-''$i''cli-multiorg1.yaml; . byfn.sh down  -c mychannel; remove_docker.sh >> /dev/null 2>&1"


ssh praiavermelha "cd ~gta/gustavo/secTry/access-control-network-5-orgs; export i=$i; docker-compose -f docker-compose-''$i''cli-multiorg2.yaml down; rm docker-compose-''$i''cli-multiorg2.yaml; . byfn.sh down -c mychannel; remove_docker.sh >> /dev/null 2>&1"


ssh flamengo "cd ~gta/gustavo/secTry/access-control-network-5-orgs; export i=$i; docker-compose -f docker-compose-''$i''cli-multiorg3.yaml down; rm docker-compose-''$i''cli-multiorg3.yaml; . byfn.sh down -c mychannel; remove_docker.sh >> /dev/null 2>&1"


ssh portuguesa "cd ~gta/gustavo/secTry/access-control-network-5-orgs; export i=$i; docker-compose -f docker-compose-''$i''cli-multiorg4.yaml down; rm docker-compose-''$i''cli-multiorg4.yaml; . byfn.sh down -c mychannel; remove_docker.sh >> /dev/null 2>&1"


ssh leblon "cd ~gta/gustavo/secTry/access-control-network-5-orgs; export i=$i; docker-compose -f docker-compose-''$i''cli-multiorg5.yaml down; rm docker-compose-''$i''cli-multiorg5.yaml; . byfn.sh down -c mychannel; remove_docker.sh >> /dev/null 2>&1"


docker-compose -f docker-compose-etcdraft2-modified.yaml down #>> /dev/null 2>&1 &
. byfn.sh down -c mychannel 
remove_docker.sh



