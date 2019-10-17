import time
import logging
from .database import *

LOGGING_FILE = "/var/log/meteor/core/core.log"
LOGGING_LEVEL = logging.DEBUG    #can be .DEBUG, .INFO, .WARNING, .ERROR, .CRITICAL

for handler in logging.root.handlers[:]:
    logging.root.removeHandler(handler)

logging.basicConfig(filename=LOGGING_FILE, filemode='w+', format='%(asctime)s - %(levelname)s - METEOR - %(message)s ', level=LOGGING_LEVEL)

def registerBot(uuid, interval, delta, hostname):
    print("registering bot")
    try:
        q = session.query(Host).filter(Host.hostname == hostname).one()
        hostid = q.id
    except:
        logstr = "METEORAPP - Bot failed to register - uuid:" + uuid + " host:" + hostname
        logging.error(logstr)
        return [False, "Unknown hostname"]
    b = Bot(uuid, interval, delta, hostid)
    logstr = "METEORAPP - Bot registered - uuid:" + uuid + " host:" + hostname
    logging.info(logstr)
    return [True, "None"]


def registerHost(hostname, interface):
    print("registering host")
    try:
        h = Host(hostname, interface)
    except:
        logstr = "METEORAPP - Host failed to register - host:" + hostname
        logging.error(logstr)
        return [False, "Host registration error"]
    logstr = "METEORAPP - Host registered - host:" + hostname
    logging.info(logstr)
    return [True, "None"]


def registerGroup(groupname):
    print("registering group")
    try:
        g = Group(groupname)
    except:
        logstr = "METEORAPP - Group failed to register - group:" + groupname
        logging.error(logstr)
        return [False, ""]
    logstr = "METEORAPP - Group registered - group:" + groupname
    logging.info(logstr)
    return [True, "None"]


def buildGroup(buildstr):
    sp = buildstr.split("||")
    for item in sp:
        if item != "":
            t = item.split(":")
            host = t[0]
            group = t[1]
            hid = session.query(Host).filter(Host.hostname == host).one().id
            gid = session.query(Group).filter(Group.name == group).one().id
            HostGroupMap(hid, gid)


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
    q = session.query(HostGroupMap).filter(HostGroupMap.groupid == groupid)
    for result in q:
        hid = result.hostid
        singlecommandadd(mode, arguments, options, hid)
    return [True, "None"]


def addGroupAction(groupname, mode, arguments, options):
    gid = grouplookup(groupname)
    if gid == "ERROR":
        return [False, "Unknown group"]
    groupcommandadd(mode, arguments, options, gid)
    return [True, "None"]

def getCommandUtil(uuid):
    t = int(time.time())
    b = session.query(Bot).filter(Bot.uuid == uuid).one()
    b.lastseen = t
    hid = b.hostid
    h = session.query(Host).filter(Host.id == hid).one()
    h.lastseen = t
    q = session.query(Action).filter(Action.hostid == hid, Action.queued == False)
    cmds = []
    for actn in q:
        aid = actn.id
        mode = actn.mode
        args = actn.arguments
        opts = actn.options
        actn.queued = True
        cmd = {"id": aid, "mode": mode, "arguments": args, "options": opts}
        cmds.append(cmd)
    session.commit()
    hstnm = h.hostname
    #updatePwnboard(hstnm)      #uncomment this to update pwnboard
    logstr = "BOXACCESS " + hstnm + " via Meteor beacon"
    logging.info(logstr)
    return cmds

def newActionResultUtil(actionid, data):
    Response(data, actionid)
    q = session.query(Action).filter(Action.id == actionid).one()
    q.responded = True
    session.commit()
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

def listGroupMembersUtil():
    data = ""
    groupstuff = {}
    for instance in session.query(HostGroupMap).order_by(HostGroupMap.id):
        hid = instance.hostid
        gid = instance.groupid
        hn = session.query(Host).filter(Host.id == hid).one().name
        gn = session.query(Group).filter(Group.id == gid).one().name
        if gn == "all":
            continue
        if gn in groupstuff:
            groupstuff[gn].append(hn)
        else:
            groupstuff[gn] = []
            groupstuff[gn].append(hn)
    return str(groupstuff)
        


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
    session.query(Response).delete()
    session.query(Action).delete()
    session.query(Bot).delete()
    session.query(Host).delete()
    session.query(Group).delete()
    session.query(HostGroupMap).delete()
    try:
        session.commit()
        logging.info("Database was cleared")
        return "Success\n"
    except:
        session.rollback()
        logging.info("Database clear failed")
        return "Error\n"

def updatePwnboard(ip):
    host = os.environ.get("PWNBOARD_URL", "")
    msg = "Meteor received a beacon"
    if not host:
        return
    data = {'ip': ip, 'application': "Meteor", 'message': msg}
    try:
        req = requests.post(host, json=data, timeout=3)
        return True
    except Exception as E:
        print("Cannot update pwnboard: {}".format(E))
        return False
