from mininet.net import Mininet
from mininet.node import Controller, RemoteController, OVSController
from mininet.node import CPULimitedHost, Host, Node
from mininet.node import OVSKernelSwitch, UserSwitch
from mininet.node import IVSSwitch
from mininet.cli import CLI
from mininet.log import setLogLevel, info
from mininet.link import TCLink, Intf, Link
from mininet.util import makeNumeric, custom
from subprocess import call

def myNetwork():
    link = custom(TCLink, bw=10)
    net = Mininet( topo=None,
                   build=False,
                   ipBase='10.0.0.0/8')

    info( '*** Adding controller\n' )
    info( '*** Add switches\n')



    s5 = net.addSwitch('s5', cls=OVSKernelSwitch, failMode='standalone')
    s2 = net.addSwitch('s2', cls=OVSKernelSwitch, failMode='standalone')
    s7 = net.addSwitch('s7', cls=OVSKernelSwitch, failMode='standalone')
    s4 = net.addSwitch('s4', cls=OVSKernelSwitch, failMode='standalone')

    r8 = net.addHost('r8', cls=Node, ip='0.0.0.0')
    r8.cmd('sysctl -w net.ipv4.ip_forward=1')
    r9 = net.addHost('r9', cls=Node, ip='0.0.0.0')
    r9.cmd('sysctl -w net.ipv4.ip_forward=1')
    r10 = net.addHost('r10', cls=Node, ip='0.0.0.0')
    r10.cmd('sysctl -w net.ipv4.ip_forward=1')

    info( '*** Add hosts\n')
    h1 = net.addHost('h1', cls=Host, ip='146.164.69.201', defaultRoute=None)
    h2 = net.addHost('h2', cls=Host, ip='146.164.69.200', defaultRoute=None)

    info( '*** Add links\n')
    net.addLink(s2, r9)
    net.addLink(s4, r10)
    net.addLink(r9, s5)
    net.addLink(s7, r8)
    net.addLink(r8, h2)
    net.addLink(s5, h2)
    net.addLink(s2, h1)
    net.addLink(s4, h1)
    net.addLink(r10, s7)

    info( '*** Starting network\n')
    net.build()
    info( '*** Starting controllers\n')
    for controller in net.controllers:
        controller.start()

    info( '*** Starting switches\n')
    net.get('s5').start([])
    net.get('s2').start([])
    net.get('s7').start([])
    net.get('s4').start([])

    info( '*** Post configure switches and hosts\n')

    CLI(net)
    net.stop()

if __name__ == '__main__':
    setLogLevel( 'info' )
    myNetwork()
