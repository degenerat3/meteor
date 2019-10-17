import os
from flask_sqlalchemy import SQLAlchemy
from daddy_app import app

db = SQLAlchemy(app)

def initDB():
    db.create_all()
    user = User(username="admin")
    user.hash_password(app.config['admin_password'])
    db.session.add(user)
    db.session.commit()

if __name__ == '__main__':
    if not os.path.exists('db.sqlite'):
        initDB()
    app.run(debug=True, host="0.0.0.0", port=8888)