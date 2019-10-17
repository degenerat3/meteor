import requests
from getpass import getpass
from requests.auth import HTTPBasicAuth

adminpass = getpass("admin password: ")
newuser = input("new username: ")
newpassword = input("new password: ")

url = "http://localhost:8888/api/users"
data = {'username': newuser, 'password': newpassword}
r = requests.post(url, auth=HTTPBasicAuth('admin', adminpass), json=data)
print(r.text)