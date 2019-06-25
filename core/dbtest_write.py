from database import *
tgroup = Group("testgroup")
thost = Host("test_hostname.com", "eth0", tgroup.id)
tbot = Bot("aaaaaaaaaaaaaaaaaa", 60, 3, thost.id)
session.add(tgroup)
session.add(thost)
session.add(tbot)
session.commit()