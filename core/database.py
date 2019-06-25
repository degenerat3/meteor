from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import Column, Integer, String, Boolean, ForeignKey
from sqlalchemy.orm import sessionmaker
Base = declarative_base()

engine = create_engine('sqlite:////tmp/test.db', echo=True)
Session = sessionmaker(bind=engine)
session = Session()

class Host(Base):
    __tablename__ = 'hosts'

    id = Column(Integer, primary_key=True)
    hostname = Column(String)
    interface = Column(String)
    groupid = Column(Integer, ForeignKey('groups.id'))

    def __repr__(self):
        return "<Host(id='%d', hostname='%s', interface='%s', groupid='%d')>" % (self.id, self.hostname, self.interface, self.groupid)

    
class Bot(Base):
    __tablename__ = 'bots'

    id = Column(Integer, primary_key=True)
    uuid = Column(String)
    interval = Column(Integer)
    delta = Column(Integer)
    hostid = Column(Integer, ForeignKey('hosts.id'))

    def __repr__(self):
        return "<Bot(id='%s', uuid='%s', interval='%d', delta='%d', hostid='%d')>" % (self.id, self.uuid, self.interval, self.delta, self.hostid)


class Group(Base):
    __tablename__ = 'groups'

    id = Column(Integer, primary_key=True)
    name = Column(String)

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
    hostid = Column(Integer, ForeignKey('host.id'))

    def __repr__(self):
        return "<Action(id='%d', mode='%s', arguments='%s', options='%s', queued='%s', responded='%s', hostid='%d')>" % (self.id, self.mode, self.arguments, self.options, self.queued, self.responded, self.hostid)
    
class Response(Base):
    __tablename__ = 'responses'

    id = Column(Integer, primary_key=True)
    data = Column(String)
    actionid = Column(Integer, ForeignKey('action.id'))

    def __repr__(self):
        return "<Respnose(id='%d', data='%s', actionid='%d')>" % (self.id, self.data, self.actionid)
