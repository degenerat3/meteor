from sqlalchemy.ext.declarative import declarative_base
Base = declarative_base()

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
        return "Bot(id='%s', uuid='%s', interval='%d', delta='%d', hostid='%d')>" % (self.id, self.uuid, self.interval, self.delta, self.hostid)


