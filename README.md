# meteor
THIS IS IN DEVELOPMENT: it is extremely unstable

It's a dockerized C2 with a flask/postgres(sqlalchemy) backend.  Modules/bots are written in golang (for now).  

**Don't be upset with the lack of any coherent documentation (or any at all for that matter)... it's coming soon<sup>TM</sup>**
## Terms
**Action:** A command or other item to be executed/handled by the client machine.  Throughout the sourcecode it is used interchangeably with the term "command."  This may change in the future to make the language more consistent.

**Client:** A program that is run on the host (of either the victim or commander) in order to send data to a module

**Core:** the internal API that interacts with the database to register bots/hosts, track actions/results,and genrally manage the C2  

**MAD:** the pseudo "protocol" that is used by certain clients to communicate with their module.  This isn't anything that has an RFC or is enforced at all, it's simply a suggestion as to how to format data when sending over protocols of a *raw* nature (TCP, UDP, ICMP, etc)

**Module:** a dockerized server that is exposed to the internet and acts as an interface between clients and the internal core

## General
The server, including the core API, the database, and all modules, is run via docker-compose.  Each module exposes port(s) to the host, so all callbacks can be directed at the same place.  The actual containers and their private network are not exposed directly to the outside world.  

The "core" of the project is made up of two portions: a Postgresql DB and a flask application.  The database stores all information related to hosts/groups/bots/actions/results.  The flask app is the API that interacts with said database, setting and retrieving values for modules.  

Separate from the core, but connected via private docker network, are the "modules."  The modules can be written in any language as long as they are dockerized and can send web requests to the core.  Actual spec for these modules can be found in the docs.

## Installation/Usage
For the server, just `docker-compose up` from the root directory of the project and you should be good to go (be careful with that because everything is in debug mode, so you'll see lots of output).  

For the client(s): they have to be slightly customized with IP/port information, then compiled for the target OS.


![The Land Before Time](docs/images/lbft.jpeg?raw=true "Image source: Hulu.com")
