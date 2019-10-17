from flask import Flask

app = Flask(__name__)
app.config['SECRET_KEY'] = 'Move fast and break things'
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///db.sqlite'
app.config['SQLALCHEMY_COMMIT_ON_TEARDOWN'] = True
app.config['admin_password'] = "breakthings"

from . import views