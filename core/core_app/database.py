import sys

from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String, Boolean, ForeignKey, create_engine, exc, MetaData
from sqlalchemy.orm import sessionmaker, scoped_session

Base = declarative_base()
engine = create_engine('postgresql://met:dbpassword@172.69.1.2/meteor', echo=True)

class Host(Base):
    __tablename__ = 'hosts'

    id = Column(Integer, primary_key=True)
    hostname = Column(String, unique=True)
    interface = Column(String)
    groupid = Column(Integer, ForeignKey('groups.id'))

    def __init__(self, hostname, interface, groupid):
        self.hostname = hostname
        self.interface = interface
        self.groupid = groupid
        session.add(self)
        try:
            session.commit()
        except exc.IntegrityError as err:
            session.rollback()
            sys.stderr.write("Error creating Host...\n")

    def __repr__(self):
        return "<Host(id='%d', hostname='%s', interface='%s', groupid='%d')>" % (self.id, self.hostname, self.interface, self.groupid)

    
class Bot(Base):
    __tablename__ = 'bots'

    id = Column(Integer, primary_key=True)
    uuid = Column(String, unique=True)
    interval = Column(Integer)
    delta = Column(Integer)
    hostid = Column(Integer, ForeignKey('hosts.id'))

    def __init__(self, uuid, interval, delta, hostid):
        self.uuid = uuid
        self.interval = interval
        self.delta = delta
        self.hostid = hostid
        session.add(self)
        try:
            session.commit()
        except exc.IntegrityError as err:
            session.rollback()
            sys.stderr.write("Error creating Bot...\n")

    def __repr__(self):
        return "<Bot(id='%s', uuid='%s', interval='%d', delta='%d', hostid='%d')>" % (self.id, self.uuid, self.interval, self.delta, self.hostid)


class Group(Base):
    __tablename__ = 'groups'

    id = Column(Integer, primary_key=True)
    name = Column(String, unique=True)

    def __init__(self, name):
        self.name = name
        session.add(self)
        try:
            session.commit()
        except exc.IntegrityError as err:
            session.rollback()
            sys.stderr.write("Error creating Group...\n")

    def __repr__(self):
        return "<Group(id='%d', name='%s')>" % (self.id, self.name)

class Action(Base):
    __tablename__ = 'actions'

    id = Column(Integer, primary_key=True)
    mode = Column(String)
    arguments = Column(String)
    options = Column(String)
    queued = Column(Boolean)
    responded = Column(Boolean)
    hostid = Column(Integer, ForeignKey('hosts.id'))

    def __init__(self, mode, arguments, options, queued, responded, hostid):
        self.mode = mode
        self.arguments = arguments
        self.options = options
        self.queued = queued
        self.responded = responded
        self.hostid = hostid
        session.add(self)
        try:
            session.commit()
        except exc.IntegrityError as err:
            session.rollback()
            sys.stderr.write("Error creating Action...\n")

    def __repr__(self):
        return "<Action(id='%d', mode='%s', arguments='%s', options='%s', queued='%s', responded='%s', hostid='%d')>" % (self.id, self.mode, self.arguments, self.options, self.queued, self.responded, self.hostid)
    
class Response(Base):
    __tablename__ = 'responses'

    id = Column(Integer, primary_key=True)
    data = Column(String)
    actionid = Column(Integer, ForeignKey('actions.id'), unique=True)

    def __init__(self, data, actionid):
        self.data = data
        self.actionid = actionid
        session.add(self)
        try:
            session.commit()
        except exc.IntegrityError as err:
            session.rollback()
            sys.stderr.write("Error creating Response...\n")

    def __repr__(self):
        return "<Response(id='%d', data='%s', actionid='%d')>" % (self.id, self.data, self.actionid)


Base.metadata.create_all(engine)
Session = scoped_session(sessionmaker(bind=engine))
session = Session()
