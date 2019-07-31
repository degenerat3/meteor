# Module Spec
 Modules can be written in any language.  The only requirement is that they are able to recieve data from a client (by pretty much any means) and send POST data to the core endpoints.  Since modules are essentially just middle-(wo)men for the clients/core, the requirements are pretty loose.  

 ### Required Endpoints for Control Modules (C2s)
 #### /register/bot

 #### /add/actionresults

 #### /get/command  
 - Note: This endpoint returns an array of commands in a json object, modules MUST be able to handle this and appropriately split/send them


 ### Required Endpoints for Command Modules (ex: Daddy Tops)

 #### /add/command/single

 #### /add/command/group

 #### /get/actionresult

 #### /list/\<table\>

 #### /register/host (optional)

 #### /register/group (optional)

