# meteor
![Meteor](docs/images/meteor_art.png?raw=true =250x250 "I clearly dont' do graphic design...")  

THIS IS IN DEVELOPMENT: it is extremely unstable

It's a dockerized C2 with a flask/postgres(sqlalchemy) backend.  Modules/bots are written in golang.  

**Don't be upset with the lack of any coherent documentation (or any at all for that matter)... it's coming soon<sup>TM</sup>**
## Terms
**Action:** A command or other item to be executed/handled by the client machine.  Throughout the sourcecode it is used interchangeably with the term "command."  This may change in the future to make the language more consistent.

**Client:** A program that is run on the host (of either the victim or commander) in order to send data to a module

**Core:** the internal API that interacts with the database to register bots/hosts, track actions/results,and genrally manage the C2  

**MAD:** the pseudo "protocol" that is used by clients to communicate with their module.  All MAD formatting and payload generation is handled behind the scenes, so don't worry about this.

**Metcli:** the golang package [found here](https://github.com/degenerat3/metcli) that is utilized by clients/modules to build and handle payloads. This helps keep the client and module source code very simple and clean, and hopefully makes development MUCH easier by limiting the number of functions that need to be implemented.

**Module:** a dockerized server that is exposed to the internet and acts as an interface between clients and the internal core

## General
The C2 server, including the core API, the database, and all modules, is run via docker-compose.  Each module exposes port(s) to the host, so all callbacks can be directed at the same place.  The actual containers and their private network are not exposed directly to the outside world.  

The "core" of the project is made up of two portions: a Postgresql DB and a flask application.  The database stores all information related to hosts/groups/bots/actions/results.  The flask app is the API that interacts with said database, setting and retrieving values for modules.  

Separate from the core, but connected via private docker network, are the "modules."  The modules are golang binaries that communicate with the core via web requests  Actual spec for these modules can be found in the docs.

## Installation/Usage
For the server, just `docker-compose up` from the root directory of the project and you should be good to go (be careful with that because everything is in debug mode, so you'll see lots of output).  

For the client(s): they have to be slightly customized with IP/port information, then compiled for the target OS.

