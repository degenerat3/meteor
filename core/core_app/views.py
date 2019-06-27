from flask import Flask, request
from . import core_app
from .utils import *

@app.route('/')
@app.route('/status')
def index():
    return "Meteor core is running.\n"


@app.route('/register', methods=['POST'])
def process_campfire():
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