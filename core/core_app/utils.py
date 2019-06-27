from .database import *

def registerBot(uuid, interval, delta, hostname):
    print("registering bot")
    q = session.query().filter(Host.hostname == hostname).one()
    hostid = q.id
    b = Bot(uuid, interval, delta, hostid)
    return [True, "None"]


def registerHost(hostname, interface, groupname):
    print("registering host")
    q = session.query().filter(Group.name == groupname).one()
    groupid = q.id
    h = Host(hostname, interface, groupid)
    return [True, "None"]


def registerGroup(groupname):
    print("registering group")
    g = Group(groupname)
    return [True, "None"]


def dumpDatabase():
    data = "HOSTS:\n"
    for instance in session.query(Host).order_by(Host.id):
        data += str(instance) + "\n"
    data += "\nBOTS:\n"
    for instance in session.query(Bot).order_by(Bot.id):
        data += str(instance) + "\n"
    data += "\nGROUPS:\n"
    for instance in session.query(Group).order_by(Group.id):
        data += (instance) + "\n"
    data += "\nACTIONS:\n"
    for instance in session.query(Action).order_by(Action.id):
        data += (instance) + "\n"
    data += "\nRESPONSES:\n"
    for instance in session.query(Response).order_by(Response.id):
        data += (instance) + "\n"
    return data