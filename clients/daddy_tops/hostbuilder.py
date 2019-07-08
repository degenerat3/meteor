# Reads a topology, populates the DB with hosts/groups
import requests
import json
import yaml
from yaml import Loader

server = "http://localhost:8888"

def registerGroup(groupname):
    header = {'Content-type': 'application/json'}
    data = {"groupname": groupname}
    requests.post(server + "/register/group", headers=header, data=json.dumps(data))

def registerHost(hostname, interface, groupname):
    header = {'Content-type': 'application/json'}
    data = {"hostname": hostname, "interface": interface, "groupname": groupname}
    requests.post(server + "/register/host", headers=header, data=json.dumps(data))

inp = "example_input.yml"

y = yaml.load(open(inp), Loader=Loader)

for key in y:
    registerGroup(key)
    for host in y[key]:
        host = host.split(":")
        hostname = host[0]
        interface = host[1]
        registerHost(hostname, interface, key)


