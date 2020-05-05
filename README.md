# Iot Data Access Control with Blockchain and SDN

## Requirements

The system runs a blockchain network in Hyperledger Fabric plataform and uses Mininet and Ryu controller to simulate a SDN environment. Install the following requirements prior running the system:

* [Hyperledger Fabric 2.0](https://hyperledger-fabric.readthedocs.io/en/release-2.0/install.html)
* [Mininet](http://mininet.org/download/#option-2-native-installation-from-source)
* Ryu Controller
* Python 2.7 

## Running the System
---------------
The system is able to execute in multiple hosts or in a single machine.

### Terminal 1
----
Change directory to start the network. Starting from the blockchain-marketplace diretory, execute:
    
    cd hyperledger-fabric/fabric-samples/access-control-network

And:

    ./byfn.sh up

### Terminal 2
-----

Starting from the blockchain-marketplace directory, start a mininet topology with a remote controller:

    sudo mn --topo single,3 --mac --controller remote --switch ovsk

### Terminal 3
-----

Starting from the blockchain-marketplace directory, start the SDN controller, by running:

    pip install -r requirements.txt
    cd controller
    ryu-manager --verbose ac-controller.py

If Hyperledger Fabric is running in a different machine, change the following lines in blockchainRemoteInteractionModule.py to the IP address and port of the computer running Hyperledger:

    IP = "146.164.69.163"
    PORT = "2346"
In the machine running Hyperledger Fabric, run:
    
    python server.py

In the machine running the SDN controller, run:

    ryu-manager --verbose ac-remote-controller.py
