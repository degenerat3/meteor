<p align="center">
  <img width="200" height="200" src="http://link.to.meteor">
</p>

# Meteor
A cross-platform C2/teamserver supporting multiple transport protocols, written in Go. 

Note: This is in development and as such is _not exactly stable_. Documentation is also lacking, but will be improved gradually over time. 



## General 
The Meteor system is split into several parts:
 - **Core**: the main "team server" element that will track actions, hosts, groups, etc. The Core runs an internal API that is "not" accessible except from listeners and other containers on the Meteornet (the docker network).
 - **Database** - A Postgres database, holding [ent](https://entgo.io/) relations that are used by the Core
 - **Client** - How a user interacts with core. The current client is Daddy Tops, but custom options can be rolled fairly easily
 - **Agents** - The actual implants that will be run on infected hosts. Agents are responsbile for communicating with their listener, pulling/executing actions, and returning their results
 - **Listeners** - The intermediary between outsiders and the Core. Each communication protocol will have it's own listener (eg one for web, one for ICMP, etc). Listeners will process agent check-ins, then send back the pending actions, eventually forwarding the results back to the Core. Daddy Tops and Nest live in the directory with the other listeners since they are hosted containers that listen on the network, but their purpose is slightly different than the listeners used for Agent communication 

 ## Installation and Usage 
 Clone the repository and build the required compose images:
```
$ git clone https://github.com/degenerat3/meteor
$ cd meteor
$ docker-compose build
<wait patiently as Golang and docker stuff happens>
$ docker-compose up
```
Meteor is now up! Note that you can remove containers from the compose file if you won't be using them, so if you only want to use the web transport there's no need to also build and run Petrie/Cera. Once your containers are running, curl `localhost:8888` to make sure the core is running.   

At this point a Daddy Tops client should be built, so follow the instructions in `meteor/docs/daddy_tops.md` to download the client and start building agents!

 ## Protobuf 
 _Nearly all_ communication in the Meteor system is done with protocol buffers. The Meteor Communication Standard (MCS) defines how to format data for the Core to process it. The listeners and agents also utilize MCS for transfer of actions and results, and Daddy Tops utilizes MCS for everything from authentication to bot and group registration. The MCS proto file can be found at `meteor/pbuf/mcs.proto`.  

 ## Current Transport Channels 
 The current agent/listener pairs are implemented, with plans for more in the future:
  - Petrie: A basic TCP socket
  - Little_Foot: Web (HTTP)
  - Cera (WIP): ICMP  

Developing additonal listeners should be simple enough if you so desire, since most of the actual Meteor functionality is abstracted away to Agent and Listener utils. For the most part, the only thing required for creation of a new listener is a reliable way to send and recieve a byte string. From there the listener utils can route the data to the appropriate Core API endpoint, and the agent utils can execute the actions and build the proper MCS payloads. There's still a lot of improvement to be made in this regard, as there's a bit of logic and protobuf parsing done outside the utils in the main functions.  

## Nest 
The Nest is used for building and compiling Meteor code (currently only agents), so you don't have to. You can use Daddy Tops (or something custom) to send the required params. Unlike the rest of the project, the Nest API is comprised of JSON endpoints rather than protobuf. This is so the binaries can be built and downloaded with a few simple curl commands, rather than requring use of protbuf and more complicatd code. 

## More Docs 
Docs are fairly limited right now, but will be improving over time. Some docs for the Core and Nest APIs, as well as instructions for Daddy Tops can be found in `meteor/docs`. 


**DISCLAIMER**: This tool is for educational purposes only. Don't diddle machines that aren't yours. The authors are not responsbile for any illicit uses of this codebase. 