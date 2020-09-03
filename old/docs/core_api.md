# Core API Endpoints

## Registration Endpoints

### /register/bot  
Description:  Register a new bot with the database  
Method: POST  
Parameters: 
```
uuid:      the client-generated unique identifier for the new bot [string]
interval:  the interval (in seconds) between callbacks [string]
delta:     the variation (in seconds) of each interval [string]
hostname:  the host that the bot is running on (generally an IP address) [string]
```

### /register/host 
Description:  Register a new host with the database  
Method: POST  
Parameters: 
```
hostname:   the host to be registered (usually an IP address) [string]
interface:  the name of the primary networking interface (ex: ens33, eth1, etc) [string]
groupname:  the name of the group that the host belongs to (web, windows, etc) [string]
```

### /register/group
Description:  Register a new group with the database  
Method: POST  
Parameters: 
```
groupname:  the name of the group being created [string]
```

## Add Endpoints
### /add/command/single
Description:  Queue an action targeting a single host
Method: POST  
Parameters: 
```
hostname:   the host that the action will be run on [string]
mode:       the type of action being executed (see mode documentation for options) [string]
arguments:  data neccessary for the action selection (ex: command line arguments) [string]
options:    for future use, not currently utilized [string]
```

### /add/command/group
Description:  Queue an action targeting an entire group  
Method: POST  
Parameters: 
```
groupname:  the group of hosts that the action will be run on [string]
mode:       the type of action being executed (see mode documentation for options) [string]
arguments:  data neccessary for the action selection (ex: command line arguments) [string]
options:    for future use, not currently utilized [string]
```

### /add/actionresult
Description:  Track feedback of action (stdout/stderr/etc)  
Method: POST  
Parameters: 
```
actionid:   the ID of the action that was executed [string] 
data:       the output of the action (stdout/stderr/etc) [string]
```
## Get Endpoints
### /get/actionresult
Description:  View the feedback from an executed action
Method: POST  
Parameters: 
```
actionid:   the ID of the action you want to view the result of [string]
```

### /get/command
Description:  Fetch all actions that are currently queue'd for your host
Method: POST  
Parameters: 
```
uuid:   the unique ID of the bot requesting the actions [string]
```

### List Endpoints
### /list/actions
Description:  show details of all registered actions  
Method: GET

### /list/bots
Description:  show details of all registered bots  
Method: GET 

### /list/groups
Description:  show details of all registered groups  
Method: GET

### /list/hosts
Description:  show details of all registered hosts  
Method: GET

## Other Endpoints
### /dumpdb or /list/db
Description:  show every record in every table of the meteor DB  
Method: GET

### /cleardb
Description:  delete every record in every table of the meteor DB  
Method: GET

