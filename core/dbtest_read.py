from database import *

for instance in session.query(Bot).order_by(Bot.id):
    print(instance)