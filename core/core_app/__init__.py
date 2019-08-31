from flask import Flask
import logging

app = Flask(__name__)

l = logging.getLogger('werkzeug')
l.disabled = True

from . import views