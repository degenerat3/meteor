"""
All the code related authentication is taken from https://github.com/miguelgrinberg/REST-auth

Huge shoutout to him for making basic token auth easy af <3
"""
from flask import Flask, request, abort, jsonify, g, url_for
from flask_sqlalchemy import SQLAlchemy
from flask_httpauth import HTTPBasicAuth
from passlib.apps import custom_app_context as pwd_context
from itsdangerous import (TimedJSONWebSignatureSerializer
                          as Serializer, BadSignature, SignatureExpired)
from . import app
import requests
import json
import datetime


core = "http://172.69.1.1:9999"
db = SQLAlchemy(app)
auth = HTTPBasicAuth()


currentDT = datetime.datetime.now()
fname = str(currentDT.strftime("%Y%m%d")) + "-event.log"
logfile = ("/var/log/meteor/daddytops/" + fname)
tmp = open(fname, "w+")
tmp.close() #init the log file


def logAction(logstr):
    ct = datetime.datetime.now()
    bigstr = str(ct.strftime("%Y-%m-%d %H:%M:%S")) + " " + logstr
    f = open(logfile, 'a')
    f.write(bigstr)
    f.close()

@app.route('/', methods=['GET'])
@app.route('/status', methods=['GET'])
def index():
    return "Daddytops is running.\n"


@app.route('/register/host', methods=['POST'])
@auth.login_required
def newhost():
    content = request.json
    try:
        hostname = content['hostname']
        interface = content['interface']
    except:
        return "Missing required fields"
    header = {'Content-type': 'application/json'}
    data = {"hostname": hostname, "interface": interface}
    req = requests.post(core + "/register/host", headers=header, data=json.dumps(data))
    logstr = "Host <" + str(hostname) + "> has been registered\n"
    logAction(logstr)
    return req.text

@app.route('/register/group', methods=['POST'])
@auth.login_required
def newgroup():
    content = request.json
    try:
        groupname = content['groupname']
    except:
        return "Missing required field"
    header = {'Content-type': 'application/json'}
    data = {"groupname": groupname}
    req = requests.post(core + "/register/group", headers=header, data=json.dumps(data))
    logstr = "Group <" + str(groupname) + "> has been registered\n"
    logAction(logstr)
    return req.text

@app.route('/register/buildgroups', methods=['POST'])
@auth.login_required
def buildgroups():
    content = request.json
    try:
        buildstr = content['buildstring']
    except:
        return "Missing required field"
    header = {'Content-type': 'application/json'}
    data = {"buildstring": buildstr}
    req = requests.post(core + "/register/buildgroups", headers=header, data = json.dumps(data))
    return req.text


@app.route('/add/command/single', methods=['POST'])
@auth.login_required
def newaction():
    content = request.json
    try:
        hostname = content['hostname']
        mode = content['mode']
        arguments = content['arguments']
        options = content['options']
        dt_user = content['dtuser']
    except:
        return "Missing required field"
    header = {'Content-type': 'application/json'}
    data = {"hostname": hostname, "mode": mode, "arguments": arguments, "options": options}
    req = requests.post(core + "/add/command/single", headers=header, data=json.dumps(data))
    logstr = "User <" + dt_user + "> queued action [mode: <" + mode + ">, " + "args: <" + arguments + ">] against host <" + hostname + ">\n"
    logAction(logstr) 
    return req.text

@app.route('/add/command/group', methods=['POST'])
@auth.login_required
def newgroupaction():
    content = request.json
    try:
        groupname = content['groupname']
        mode = content['mode']
        arguments = content['arguments']
        options = content['options']
        dt_user = content['dtuser']
    except:
        return "Missing required field"
    header = {'Content-type': 'application/json'}
    data = {"groupname": groupname, "mode": mode, "arguments": arguments, "options": options}
    req = requests.post(core + "/add/command/group", headers=header, data=json.dumps(data))
    logstr = "User <" + dt_user + "> queued action [mode: <" + mode + ">, " + "args: <" + arguments + ">] against group <" + groupname + ">\n"
    logAction(logstr) 
    return req.text


@app.route('/get/actionresult', methods=['POST'])
@auth.login_required
def getactionresult():
    content = request.json
    try:
        aid = content['actionid']
    except:
        return "Missing required field"
    header = {'Content-type': 'application/json'}
    data = {"actionid": aid}
    req = requests.post(core + "/get/actionresult", headers=header, data=json.dumps(data))
    return req.text

@app.route('/list/bots', methods=['GET'])
@auth.login_required
def listbots():
    data = requests.get(core + "/list/bots")
    return data.text

@app.route('/list/hosts', methods=['GET'])
@auth.login_required
def listhosts():
    data = requests.get(core + "/list/hosts")
    return data.text

@app.route('/list/groups', methods=['GET'])
@auth.login_required
def listgroups():
    data = requests.get(core + "/list/groups")
    return data.text

@app.route('/list/actions', methods=['GET'])
@auth.login_required
def listactions():
    data = requests.get(core + "/list/actions")
    return data.text

@app.route('/list/db', methods=['GET'])
@app.route('/dumpdb', methods=['GET'])
@auth.login_required
def dumpdb():
    data = requests.get(core + "/dumpdb")
    return data.text

class User(db.Model):
    __tablename__ = 'users'
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(32), index=True)
    password_hash = db.Column(db.String(64))

    def hash_password(self, password):
        self.password_hash = pwd_context.encrypt(password)

    def verify_password(self, password):
        return pwd_context.verify(password, self.password_hash)

    def generate_auth_token(self, expiration=600):
        s = Serializer(app.config['SECRET_KEY'], expires_in=expiration)
        return s.dumps({'id': self.id})

    @staticmethod
    def verify_auth_token(token):
        s = Serializer(app.config['SECRET_KEY'])
        try:
            data = s.loads(token)
        except SignatureExpired:
            return None    # valid token, but expired
        except BadSignature:
            return None    # invalid token
        user = User.query.get(data['id'])
        return user


@auth.verify_password
def verify_password(username_or_token, password):
    # first try to authenticate by token
    user = User.verify_auth_token(username_or_token)
    if not user:
        # try to authenticate with username/password
        user = User.query.filter_by(username=username_or_token).first()
        if not user or not user.verify_password(password):
            return False
    g.user = user
    return True


@app.route('/api/users', methods=['POST'])
@auth.login_required
def new_user():
    username = request.json.get('username')
    password = request.json.get('password')
    if username is None or password is None:
        abort(400)    # missing arguments
    if User.query.filter_by(username=username).first() is not None:
        abort(400)    # existing user
    user = User(username=username)
    user.hash_password(password)
    db.session.add(user)
    db.session.commit()
    return (jsonify({'username': user.username}), 201,
            {'Location': url_for('get_user', id=user.id, _external=True)})


@app.route('/api/users/<int:id>')
def get_user(id):
    user = User.query.get(id)
    if not user:
        abort(400)
    return jsonify({'username': user.username})


@app.route('/api/token')
@auth.login_required
def get_auth_token():
    token = g.user.generate_auth_token(600)
    return jsonify({'token': token.decode('ascii'), 'duration': 6000})

@app.route('/api/testauth')
@auth.login_required
def get_resource():
    return jsonify({'data': 'Hello, %s!' % g.user.username})

def initDB():
    db.create_all()
    user = User(username="admin")
    user.hash_password(app.config['admin_password'])
    db.session.add(user)
    db.session.commit()