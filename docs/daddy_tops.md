# Daddy Tops Client Guide
A quick reference for using Daddy Tops to interact with Meteor   

## General 
The Daddy Tops utility serves as the main avenue for interacting with the Meteor infrastructure. It has an interactive prompt for interacting with bots/actions/etc, as well as several command-line argument options for registering new hosts, user management, and setting the upstream DT server.   

## Acquiring the Client
When the docker containers are brought up, the latest DT will be compiled and hosted on via the Nest. They can be downloaded from the following locations:
```
Linux:  [nest_root]/files/dt/nix_dt.bin
Win:    [nest_root]/files/dt/win_dt.exe
```

Only DT versions for Linux and Windows automatically compiled, others can be built from the source located in `meteor/client/daddy_tops`.   

## Command-Line Options
 - `./dt --server <server_ip>:<port>` - will set the appropriate env variable to save the upstream server location  
 - `./dt --register-hosts <config.yml>` - parse a hosts yml file, register the host/group objects in the meteor backend. An example config file can be found at `meteor/client/daddy_tops/example_hosts.yml`  
 - `./dt --register-user` - follow the prompts to add a new user. This requires the admin password
 - `./dt --change-password` - follow prompts to update the password for an existing target user. Can be done by admin or the target user themself
 - `./dt --clear-token` - clear your session token so you can log in as a different user

## Interactive Prompt  
In order to access the interactive CLI, simply run `./dt` without any arguments and follow the login prompt. Afterwhich you will be dumped to the Daddy Tops shell: `DT>`. From here, the user has access to the following utilities:

```
CAPABILITY				SYNTAX
------------------------------------------------------------------------------------------

NEW ACTION:             action <%target_hostname%> <%mode_code%> <%arguments%>
NEW GROUP ACTION:       gaction <%target_groupname%> <%mode_code%> <%arguments%>
SHOW RESULT:            result <%uuid%>
LIST AVAILABLE <X>:     list <modes/hosts/host/groups/group/bots/actions> <OPT:%host%/%group%>
BUILD REQUEST:          build <agent/*>             
HELP MENU               help
QUIT PROMPT             exit
```

The "mode code" used for action creation must come from this list, more options coming in the future:

```
MODE    DESC                ARGS	
-------------------------------------
  1     shell exec          <shell command>
  2     firewall flush      N/A
  3     create priv user    <username>
  4     enable SSH/RDP      N/A
  F     nuke the box        N/A
```

So an example to queue up a new action against the host `8.8.8.8` would look like:
 ```
 action 8.8.8.8 1 cat /etc/passwd
 ```
In the above, the "shell" mode (`1`) is being used to cat the contents of a file, which can be later accessed by using the `result` command. Sometimes you may want to execute an action against an entire group of hosts. For that, use `gaction`:
```
gaction t1desktops 1 dir C:\Windows\ 
```
The above command will queue up a separate action for each host in the `t1desktops` group. Be careful when using shell exec group actions. Since the cmdlet will be run against all hosts in the group, if you pass a PowerShell command (used by Windows agents) to a group containing linux hosts, it will still try to run. While the errors shouldn't be fatal, it will fail to execute on those hosts.   

The `result` option is used to view the output of your actions. Simply pass in the action UUID as an argument and your client will display the collected output.   

The `list` option is used to view objects that are currently being tracked in the Meteor backend. Ex: use it to list all the currently tracked hosts/bots, all the actions you've queued, or all the available action modes.   

The `build` option will take you through an interactive prompt in order to build various meteor objects (currently only agent buildig is supported). You'll choose the agent type you want to build, target OS, beacon time, and a few other required configs. The output of the `build` command will be a Nest link to download the object. While agents are the only thing currently able to be compiled using DT `build`, future support for things such as droppers and other resources is planned.  



<p align="center">
  <img src="https://raw.githubusercontent.com/degenerat3/meteor/master/docs/images/dt_walkthrough.gif">
</p>