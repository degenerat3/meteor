from .database import *

def registerBot(uuid, interval, delta, hostname):
    print("registering bot")
    try:
        q = session.query(Host).filter(Host.hostname == hostname).one()
        hostid = q.id
    except:
        return [False, "Unknown hostname"]
    b = Bot(uuid, interval, delta, hostid)
    return [True, "None"]


def registerHost(hostname, interface, groupname):
    print("registering host")
    try:
        q = session.query(Group).filter(Group.name == groupname).one()
        groupid = q.id
    except:
        return [False, "Unknown group"]
    h = Host(hostname, interface, groupid)
    return [True, "None"]


def registerGroup(groupname):
    print("registering group")
    g = Group(groupname)
    return [True, "None"]


def hostlookup(hostname):
    try:
        q = session.query(Host).filter(Host.hostname == hostname).one()
        hostid = q.id
        return hostid
    except:
        return "ERROR"


def grouplookup(groupname):
    try:
        q = session.query(Group).filter(Group.name == groupname).one()
        gid = q.id
        return gid
    except:
        return "ERROR"


def singlecommandadd(mode, arguments, options, hostid):
    a = Action(mode, arguments, options, False, False, hostid)

def groupcommandadd(mode, arguments, options, groupid):
    q = session.query(Host).filter(Host.groupid == groupid)
    for result in q:
        hid = result.id
        singlecommandadd(mode, arguments, options, hid)
    return [True, "None"]


def addGroupAction(groupname, mode, arguments, options):
    gid = grouplookup(groupname)
    if gid == "Error":
        return [False, "Unknown host"]
    groupcommandadd(mode, arguments, options, gid)

def getCommandUtil(hostname):
    hid = hostlookup(hostname)
    q = session.query(Action).filter(Action.hostid == hid)
    cmds = []
    for actn in q:
        mode = actn.mode
        args = actn.arguments
        opts = actn.options
        actn.queued = True
        cmd = {"mode": mode, "arguments": args, "options": opts}
        cmds.append(cmd)
    session.commit()
    return cmds

def newActionResultUtil(actionid, data):
    Response(data, actionid)
    return "Success"

def getActionResultUtil(actionid):
    q = session.query(Response).filter(Response.actionid == actionid).one()
    return q.data

def listHostsUtil():
    data = ""
    for instance in session.query(Host).order_by(Host.id):
        data += str(instance) + "\n"
    return data

def listBotsUtil():
    data = ""
    for instance in session.query(Bot).order_by(Bot.id):
        data += str(instance) + "\n"
    return data

def listGroupsUtil():
    data = ""
    for instance in session.query(Group).order_by(Group.id):
        data += str(instance) + "\n"
    return data

def listActionsUtil():
    data = ""
    for instance in session.query(Action).order_by(Action.id):
        data += str(instance) + "\n"
    return data

def dumpDatabase():
    data = "HOSTS:\n"
    for instance in session.query(Host).order_by(Host.id):
        data += str(instance) + "\n"
    data += "\nBOTS:\n"
    for instance in session.query(Bot).order_by(Bot.id):
        data += str(instance) + "\n"
    data += "\nGROUPS:\n"
    for instance in session.query(Group).order_by(Group.id):
        data += str(instance) + "\n"
    data += "\nACTIONS:\n"
    for instance in session.query(Action).order_by(Action.id):
        data += str(instance) + "\n"
    data += "\nRESPONSES:\n"
    for instance in session.query(Response).order_by(Response.id):
        data += str(instance) + "\n"
    return data


def clearDbUtil():
    meta = db.metadata
    for table in reversed(meta.sorted_tables):
        session.execute(table.delete())
    try:
        session.commit()
    except:
        return "Failure"
    return "Success"