#!/bin/bash

## Usage: ./generate_certs.sh <#orgs> <#peers per org>

NUM_ORGS=$1
PEERS_PER_ORG=$2

if [ $# -lt 2 ]; then
    echo "Usage: ./generate_certs.sh <#orgs> <#peers per org>"
    exit 1
fi

echo "\
OrdererOrgs:
  - Name: Orderer
    Domain: example.com
    Specs:
      - Hostname: orderer
      - Hostname: orderer2
      - Hostname: orderer3
      - Hostname: orderer4
      - Hostname: orderer5 
PeerOrgs: \
 " > crypto-config.yaml

 for i in $(seq 1 $(($NUM_ORGS))); do
    echo "\
  - Name: Org$i
    Domain: org$i.example.com
    EnableNodeOUs: true
    Template:
      Count: $PEERS_PER_ORG
    Users:
      Count: 1 \
     " >> crypto-config.yaml
done

echo "Generated certificates for $NUM_ORGS organizations with $PEERS_PER_ORG peer(s) each."
