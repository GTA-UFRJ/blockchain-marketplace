# -*-  coding: utf-8 -*-

import subprocess


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
    lastQuote = 0
    result = []
    string = transaction
    while lastQuote != -1:
        firstQuote = string.find("\"")
        lastQuote = string[firstQuote+1:].find("\"")
        result.append(string[firstQuote+1:lastQuote+1])
        string = string[lastQuote+1:]
        if lastQuote == -1:
            end = 1
    return result

# Receives a transaction and return the query of the blockchain
def getTransaction(transaction):
    command = "docker exec -it cli peer chaincode query -C mychannel -n mycc -c \'{\"Args\":[\"getHistoryForTransaction\",\""+transaction+"\"]}\'"
    getQuery = subprocess.Popen(command,shell=True,stdout=subprocess.PIPE).stdout
    return str(getQuery.read().decode())

# transaction: the transaction that have a new IP
# dictionary: controller IP dictionary
def updateDictionary(dictionary,transaction):
    ipSellerIndex = 13
    ipBuyerIndex = 13
    typeIndex = 17
    transactionQuery = getTransaction(transaction)
    transactionFields = splitTransaction(transactionQuery)
    if transactionFields[typeIndex] == "sell":
        addKey(transactionFields[ipSellerIndex],dictionary)
    else:
        addValue(transactionFields[ipSellerIndex],transactionFields[ipBuyerIndex],dictionary)



if __name__ == "__main__":
    a = {}
    print("Before add a key and a value")
    print(a.values())
    print (a.keys())
    updateDictionary(a,"t14")
    updateDictionary(a,"t114")
    print("After add a key and a value")
    print(a.values())
    print(a.keys())
