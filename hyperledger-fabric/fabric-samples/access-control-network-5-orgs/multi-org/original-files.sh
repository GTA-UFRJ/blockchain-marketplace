#!/bin/bash
#
#
# Restore the original files of the network
#


rm ../byfn.sh
mv ../byfn.sh.original ../byfn.sh
rm ../ccp-generate.sh
mv ../ccp-generate.sh.original ../ccp-generate.sh
rm ../configtx.yaml 
mv ../configtx.yaml.original ../configtx.yaml
rm ../crypto-config.yaml 
mv ../crypto-config.yaml.original ../crypto-config.yaml
rm ../base/docker-compose-base.yaml 
mv ../base/docker-compose-base.yaml.original ../base/docker-compose-base.yaml
rm ../scripts/script.sh 
mv ../scripts/script.sh.original ../scripts/script.sh
rm ../scripts/utils.sh 
mv ../scripts/utils.sh.original ../scripts/utils.sh
