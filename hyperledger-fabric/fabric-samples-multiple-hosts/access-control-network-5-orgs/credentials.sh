#!/bin/bash

. byfn.sh generate #>> /dev/null 2>&1
scp -r crypto-config ipanema:~gta/gustavo/secTry/access-control-network-5-orgs #>> /dev/null 2>&1
scp -r channel-artifacts ipanema:~gta/gustavo/secTry/access-control-network-5-orgs #>> /dev/null 2>&1
scp -r crypto-config praiavermelha:~gta/gustavo/secTry/access-control-network-5-orgs #>> /dev/null 2>&1
scp -r channel-artifacts praiavermelha:~gta/gustavo/secTry/access-control-network-5-orgs #>> /dev/null 2>&1
scp -r crypto-config flamengo:~gta/gustavo/secTry/access-control-network-5-orgs #>> /dev/null 2>&1
scp -r channel-artifacts flamengo:~gta/gustavo/secTry/access-control-network-5-orgs #>> /dev/null 2>&1
scp -r crypto-config portuguesa:~gta/gustavo/secTry/access-control-network-5-orgs #>> /dev/null 2>&1
scp -r channel-artifacts portuguesa:~gta/gustavo/secTry/access-control-network-5-orgs #>> /dev/null 2>&1
scp -r crypto-config leblon:~gta/gustavo/secTry/access-control-network-5-orgs #>> /dev/null 2>&1
scp -r channel-artifacts leblon:~gta/gustavo/secTry/access-control-network-5-orgs #>> /dev/null 2>&1
