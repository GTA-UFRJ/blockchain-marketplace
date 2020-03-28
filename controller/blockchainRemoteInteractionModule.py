# -*-  coding: utf-8 -*-

import subprocess
import rpyc
import sys

IP = "146.164.69.163"
PORT = "2346"

# Try to add a new key in a dictionary object
# Return 0 if no error was found, erro number otherwise
# Key (string) - ex: "198.162.0.1"
#
# Erros
# 1 - key is not a string object
# 2 - key already existis
def addKey(key,dictionary):
    if type(key) != str:
        return 1
    if key not in dictionary.keys():
        dictionary[key] = []
        return 0
    return 2


# Try to add a new value in a key of a dictionary object
# Return 0 if no error was found, erro number otherwise
# key (string) - ex: "198.162.0.1"
# value (string) - ex: "10.0.0.1"
#
# Erros
# 1 - key or value is not a string object
# 2 - the value already belongs to the key
# 3 - key not in dictionary
def addValue(key,value,dictionary):
    if type(value) != str or type(key) != str:
        return 1
    if key in dictionary.keys():
        if value in dictionary[key]:
            return 2
        dictionary[key].append(value)
        return 0
    return 3

# Split the transaction and return a list of values
# transaction (string) - ex: "\"id","ip\""
# result (list) - ex: ["id",]
def splitTransaction(transaction):
    if type(transaction) != str:
        return 1
    countError = transaction.count("failure")
    if countError:
        return "", 0    
    count = transaction.count("TxId")
    #print ("\n\nCount: " + str(count))
    parsed_json = (json.loads(transaction))
    #print(json.dumps(parsed_json, indent=4, sort_keys=True))
    return parsed_json, count

# Receives a transaction and return the query of the blockchain
def getTransaction(blockchainServer):
	serverResponse = blockchainServer.root.queryTransaction() 
	return serverResponse

# Open a connection with a remote server
def openConnection():
    global IP
    global PORT
    connection = rpyc.connect(IP,PORT)
    return connection

# transaction: the transaction that have a new IP
# dictionary: controller IP dictionary
def updateDictionary(dictionary, blockchainServer):
    index = 0
    transactionQuery = getTransaction(blockchainServer)
    #print ("Print da transacao que o pyhton pegou: " +transactionQuery)
    transactionFields, count = splitTransaction(transactionQuery)
    while (count != 0):
        #print ("SrcIPAddress: " + transactionFields[index]["SrcIPAddress"])
        returnValue = addKey(str(transactionFields[index]["DstIPAddress"]),dictionary)
        returnValue = addKey(str(transactionFields[index]["SrcIPAddress"]),dictionary)
        #print ("Return Value:" + str(returnValue))
        returnValue = addValue(str(transactionFields[index]["DstIPAddress"]),str(transactionFields[index]["SrcIPAddress"]),dictionary)
        returnValue = addValue(str(transactionFields[index]["SrcIPAddress"]),str(transactionFields[index]["DstIPAddress"]),dictionary)
        #print ("Return Value:" + str(returnValue))
        count -= 1
        index += 1        
    return dictionary



if __name__ == "__main__":
    a = {}
#    if len(sys.argv) == 3:
#            IP = sys.argv[2]
#            PORT = "2346"
#    elif len(sys.argv) == 4:	
#            IP = sys.argv[2]
#            PORT = sys.argv[3]
#    else:
#            IP = "146.164.69.163"
#            PORT = "2346"
    blockchainServer = openConnection()
    updateDictionary(a,"t14", blockchainServer)
    updateDictionary(a,"t114", blockchainServer)
    print("After add a key and a value")
    print(a.values())
    print(a.keys())
