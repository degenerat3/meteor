import os
from daddy_app import app

if __name__ == '__main__':
    if not os.path.exists('db.sqlite'):
        db.create_all()
    app.run(debug=True, host="0.0.0.0", port=8888)