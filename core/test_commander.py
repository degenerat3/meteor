import requests

server = "http://localhost:9999"

header = {'Content-type': 'application/json'}
data = {groupname: "Webservers"}
request = requests.post(server + "/register/group", headers=header, data=json.dumps(data))
print(request.body)

header = {'Content-type': 'application/json'}
data = {hostname: "192.69.4.20", interface: "eth69", groupname: "Webservers"}
request = requests.post(server + "/register/host", headers=header, data=json.dumps(data))
print(request.body)

header = {'Content-type': 'application/json'}
data = {uuid: "aklsjflasdjfl", interval: 60, delta: 5, hostname: "192.69.4.20"}
request = requests.post(server + "/register/bot", headers=header, data=json.dumps(data))
print(request.body)

request = requests.get(server + "/dumpdb")