from flask import Flask, request
from . import core_app
from .utils import *

@app.route('/')
@app.route('/status')
def index():
    return "Meteor core is running.\n"


@app.route('/register/bot', methods=['POST'])
def newbot():
    content = request.json
    uuid = content['uuid']
    interval = content['interval']
    delta = content['delta']
    hostname = content['hostname']
    reg_status = registerBot(uuid, interval, delta, hostname)
    if reg_status[0]:
        return "Success"
    failure_str = "Register failure: " + reg_status[1]
    return failure_str


@app.route('/register/host', methods=['POST'])
def newhost():
    content = request.json
    hostname = content['hostname']
    interface = content['interface']
    groupname = content['groupname']
    reg_status = registerHost(hostname, interface, groupname)
    if reg_status[0]:
        return "Success"
    failure_str = "Register failure: " + reg_status[1]
    return failure_str


@app.route('/register/group', methods=['POST'])
def newgroup():
    content = request.json
    groupname = content['groupname']
    if reg_status[0]:
        return "Success"
    failure_str = "Register failure: " + reg_status[1]
    return failure_str