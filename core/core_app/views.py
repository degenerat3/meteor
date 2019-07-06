from flask import Flask, request
from . import app
from .utils import *

@app.route('/', methods=['GET'])
@app.route('/status', methods=['GET'])
def index():
    return "Meteor core is running.\n"


@app.route('/register/bot', methods=['POST'])
def newbot():
    content = request.json
    try:
        uuid = content['uuid']
        interval = content['interval']
        delta = content['delta']
        hostname = content['hostname']
    except:
        return "Missing required fields"
    reg_status = registerBot(uuid, interval, delta, hostname)
    if reg_status[0]:
        return "Success"
    failure_str = "500: Register failure- " + reg_status[1]
    return failure_str


@app.route('/register/host', methods=['POST'])
def newhost():
    content = request.json
    try:
        hostname = content['hostname']
        interface = content['interface']
        groupname = content['groupname']
    except:
        return "Missing required fields"
    reg_status = registerHost(hostname, interface, groupname)
    if reg_status[0]:
        return "Success"
    failure_str = "500: Register failure- " + reg_status[1]
    return failure_str


@app.route('/register/group', methods=['POST'])
def newgroup():
    content = request.json
    try:
        groupname = content['groupname']
    except:
        return "Missing required field"
    reg_status = registerGroup(groupname)
    if reg_status[0]:
        return "Success"
    failure_str = "500: Register failure- " + reg_status[1]
    return failure_str


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
    hostid = hostlookup(hostname)
    if hostid == "ERROR":
        return "Unknown host"
    singlecommandadd(mode, arguments, options, hostid)
    return "success"

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
    res = addGroupAction(groupname, mode, arguments, options)
    
    return "success"


@app.route('/add/actionresult', methods=['POST'])
def newactionres():
    content = request.json
    try:
        aid = content['actionid']
        data = content['data']
    except:
        return "Missing required field"
    res = newActionResultUtil(aid, data)
    return "success"

@app.route('/get/command', methods=['POST'])
def getcommand():
    content = request.json
    try:
        hostname = content['hostname']
    except:
        return "Missing required field"
    cmds = getCommandUtil(hostname)
    return cmds

@app.route('/get/actionresult', methods=['POST'])
def getactionresult():
    content = request.json
    try:
        aid = content['actionid']
    except:
        return "Missing required field"
    res = getActionResultUtil(aid)
    return "success"

@app.route('/list/bots', methods=['GET'])
def listbots():
    data = listBotsUtil()
    return data

@app.route('/list/hosts', methods=['GET'])
def listhosts():
    data = listHostsUtil()
    return data

@app.route('/list/groups', methods=['GET'])
def listgroups():
    data = listGroupsUtil()
    return data

@app.route('/list/actions', methods=['GET'])
def listactions():
    data = listActionsUtil()
    return data

@app.route('/dumpdb', methods=['GET'])
def dumpdb():
    data = dumpDatabase()
    return data