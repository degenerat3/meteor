import os
from flask_sqlalchemy import SQLAlchemy
from daddy_app import app
from views import initDB

if __name__ == '__main__':
    if not os.path.exists('db.sqlite'):
        initDB()
    app.run(debug=True, host="0.0.0.0", port=8888)