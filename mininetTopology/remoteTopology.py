from mininet.topo import Topo
#from mininet.net import Mininet
#from mininet.node import Controller, RemoteController, OVSController
#from mininet.node import CPULimitedHost, Host, Node
#from mininet.node import OVSKernelSwitch, UserSwitch
#from mininet.node import IVSSwitch
#from mininet.cli import CLI
#from mininet.log import setLogLevel, info
#from mininet.link import TCLink, Intf, Link
#from mininet.util import makeNumeric, custom
#from subprocess import call



class MyTopo(Topo):
    def __init__(self):
        # Initialize topology
        Topo.__init__(self)

        # Add hosts
        Host1 = self.addHost('h1')
        Host2 = self.addHost('h2')
        
        # Add switch
        Switch1 = self.addSwitch('s1')


        # Add links
        self.addLink(Host1,Switch1)
        self.addLink(Host2,Switch1)

topos = {'mytopo' : (lambda: MyTopo())}
       # s = self.addSwitch('s1')
       # h1 = self.addHost('h1',defaultRoute='via 146.164.69.2')
       # h2 = self.addHost('h2',defaultRoute='via 146.164.69.2')
       # self.addLink(h1,s)
       # self.addLink(h2,s)
