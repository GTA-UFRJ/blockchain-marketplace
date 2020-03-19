import subprocess
import threading
from rpyc.utils.server import ThreadedServer
import rpyc


PORT = 2346

class controllerService(rpyc.Service):
	def on_connect(self,arg):
		pass
	def on_disconnect(self,arg):
		pass
	def exposed_queryTransaction(self, transaction):
   	        command = "docker exec -it cli peer chaincode query -C mychannel -n mycc -c \'{\"Args\":[\"getHistoryForTransaction\",\""+transaction+"\"]}\'"
	        getQuery = subprocess.Popen(command,shell=True,stdout=subprocess.PIPE).stdout
	        return str(getQuery.read().decode())

def main():
	global PORT
	server = ThreadedServer(controllerService, port=PORT)
	threading.Thread(target=server.start).start()

if __name__ == "__main__":
	main()
