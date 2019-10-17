#!/usr/bin/env python3
# Reads a topology, populates the DB with hosts/groups
import requests
import json
import os
import sys
import yaml
from yaml import Loader
from requests.auth import HTTPBasicAuth
from getpass import getpass

server = os.environ.get("DT_SERVER", "http://localhost:8888") 

adminpass = getpass("admin password: ")

def registerGroup(groupname):
    header = {'Content-type': 'application/json'}
    data = {"groupname": groupname}
    requests.post(server + "/register/group", headers=header, data=json.dumps(data), auth=HTTPBasicAuth('admin', adminpass))

def registerHost(hostname, interface):
    header = {'Content-type': 'application/json'}
    data = {"hostname": hostname, "interface": interface}
    requests.post(server + "/register/host", headers=header, data=json.dumps(data), auth=HTTPBasicAuth('admin', adminpass))

def buildGroups(buildstr):
    header = {'Content-type': 'application/json'}
    data = {"buildstring": buildstr}
    requests.post(server + "/register/buildgroups", headers=header, data=json.dumps(data), auth=HTTPBasicAuth('admin', adminpass))

inp = "example_input.yml"

if len(sys.argv) > 1:
    inp = sys.argv[1]
else:
    inp = input("Hosts.yml file: ")

print("[+] Parsing yml...")
y = yaml.load(open(inp), Loader=Loader)

buildgroupstr = ""
print("[+] Registering hosts...")
for key in y:
    if key == "all":
        for host in y[key]:
            host = host.split(":")
            hostname = host[0]
            interface = host[1]
            registerHost(hostname, interface)

print("[+] Building groups...")
for key in y:
    registerGroup(key)
    for host in y[key]:
        host = host.split(":")[0]
        tmp = host + ":" + key + "||"
        buildgroupstr += tmp

buildGroups(buildgroupstr)


