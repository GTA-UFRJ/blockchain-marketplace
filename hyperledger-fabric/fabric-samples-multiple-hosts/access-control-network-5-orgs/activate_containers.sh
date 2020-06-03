#!/bin/bash

i=$1        
        

ssh ipanema "cd ~gta/gustavo/secTry/access-control-network-5-orgs; export i=$i; cp docker-multiple-''$i''/docker-compose-''$i''cli-multiorg1.yaml .; docker-compose -f docker-compose-''$i''cli-multiorg1.yaml up -d >> /dev/null 2>&1"


ssh praiavermelha "cd ~gta/gustavo/secTry/access-control-network-5-orgs; export i=$i; cp docker-multiple-''$i''/docker-compose-''$i''cli-multiorg2.yaml .; docker-compose -f docker-compose-''$i''cli-multiorg2.yaml up -d >> /dev/null 2>&1"


ssh flamengo "cd ~gta/gustavo/secTry/access-control-network-5-orgs; export i=$i; cp docker-multiple-''$i''/docker-compose-''$i''cli-multiorg3.yaml .; docker-compose -f docker-compose-''$i''cli-multiorg3.yaml up -d >> /dev/null 2>&1"


ssh portuguesa "cd ~gta/gustavo/secTry/access-control-network-5-orgs; export i=$i; cp docker-multiple-''$i''/docker-compose-''$i''cli-multiorg4.yaml .; docker-compose -f docker-compose-''$i''cli-multiorg4.yaml up -d >> /dev/null 2>&1"


ssh leblon "cd ~gta/gustavo/secTry/access-control-network-5-orgs; export i=$i; cp docker-multiple-''$i''/docker-compose-''$i''cli-multiorg5.yaml .; docker-compose -f docker-compose-''$i''cli-multiorg5.yaml up -d >> /dev/null 2>&1"


docker-compose -f docker-compose-etcdraft2-modified.yaml up -d >> /dev/null 2>&1 &


