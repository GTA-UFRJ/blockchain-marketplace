# Access-Control-SDN

## Controller and blockchain in the same machine

If the blockchain is running in the same machine that RYU, use the code ac-controller.py

## Controller in a remote machine

If the blockchain is running in a different machine that RYU, use the follow procedures:
* Run the command "pip install -r requirements.txt"
* Execute the server.py code in the machine that runs the blockchain
* Execute the ac-controller-remote.py in the SDN machine
