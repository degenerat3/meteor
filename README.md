# meteor
![Meteor](docs/images/small_meteor_art.png?raw=true "I clearly don't do graphic design...")  

THIS IS IN DEVELOPMENT: it is extremely unstable

It's a dockerized C2 with a flask/postgres(sqlalchemy) backend.  Modules/bots are written in golang.  

**Don't be upset with the lack of any coherent documentation (or any at all for that matter)... it's coming soon<sup>TM</sup>**
## Terms
**Action:** A command or other item to be executed/handled by the client machine.  Throughout the sourcecode it is used interchangeably with the term "command."  This may change in the future to make the language more consistent.

**Client:** A program that is run on the host (of either the victim or commander) in order to send data to a module

**Core:** the internal API that interacts with the database to register bots/hosts, track actions/results,and genrally manage the C2  

**MAD:** the pseudo "protocol" that is used by clients to communicate with their module.  All MAD formatting and payload generation is handled behind the scenes, so don't worry about this.

**Metcli:** the golang package [found here](https://github.com/degenerat3/metcli) that is utilized by clients/modules to build and handle payloads. This helps keep the client and module source code very simple and clean, and hopefully makes development MUCH easier by limiting the number of functions that need to be implemented.  

**Modes:** the modes are the "opcodes" used by the clients for action execution. For example, the mode "1" is used for shell execution, mode "2" is used to flush firewalls. Currently-supported opcodes are listed in docs/MAD.md.  

**Module:** a dockerized server that is exposed to the internet and acts as an interface between clients and the internal core

## General
The C2 server, including the core API, the database, and all modules, is run via docker-compose.  Each module exposes port(s) to the host, so all callbacks can be directed at the same place.  The actual containers and their private network are not exposed directly to the outside world.  

The "core" of the project is made up of two portions: a Postgresql DB and a flask application.  The database stores all information related to hosts/groups/bots/actions/results.  The flask app is the API that interacts with said database, setting and retrieving values for modules.  

Separate from the core, but connected via private docker network, are the "modules."  The modules are golang binaries that communicate with the core via web requests  Actual spec for these modules can be found in the docs.

## Installation  
#### Server   
First, clone the repository: `git clone https://github.com/degenerat3/meteor`  
Ensure that you have [docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/install/) installed.  
Move into the root of the `meteor` directory.  
Build the containers: `sudo docker-compose build`  
Start (and daemonize) the containers: `sudo docker-compose up -d`  
To test if core is running, from localhost: `curl http://172.69.1.1:9999/`  

#### Clients
Each client must be slightly adjusted, since there are several global variables that must be configred.  
The compilation can be done automatically with the `clients/build.sh` script (run it, input prompted data, ezwin).  
If you want to do it yourself:  
The following variables must be configured (names may vary slightly by client):  
 - SERV     // the server to call back to (docker host)
 - REGFILE  // The destination file for registration info (obfuscated UUID)
 - OBFSEED  // The seed integer for the registration obfuscation
 - OBFTEXT  // The seed text used for registration obfuscation  

 Once the the variables are configured, the binary must be built.  
 Set the golang env for target OS: `set goos=linux` or whatever, see [this](https://golang.org/pkg/go/build/) for more info on golang build options.  
 Running `go build` in the client directory will generate the proper .exe or .elf file (depending on target). 
 The compiled binary can now be run on the target victim machines, usage instructions vary per bot.  



## Usage
The general flow is that a commander module (the default is `DaddyTops`) will be used to set up hosts/groups in the Core (only done once), then a user interacts with the command module in order to queue commands for hosts, view results, etc. Instructions on how to use the DaddyTops module can be found in `docs/daddytops.md`.