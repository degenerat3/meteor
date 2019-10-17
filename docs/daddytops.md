# DaddyTops
The default Commander module for meteor

## Installation/Setup  
In the base directory of the module/client, run `pip3 install -r requirements.txt` in order to fetch the neccessary python packages.

#### User Creation
The first step after the database is brought up is to create users for DaddyTops. There is a default `admin` user, who's password is set in the daddy_app `__init__.py` file. To create another user, simply run the `user_creator.py` and enter the admin password when prompted. It will ask for a new username and new password, and when supplied it will create the new user, which can now be used to log into the CLI.

#### Host/Group Construction
Next, before you are able to control any bots or issue any commands, is to build the hosts/groups into the DB. There is a tool for this located at `modules/daddy_tops/utils/hostbuilder.py`. The input for `hostbuilder.py` is a `.yml` file containing the hosts and groups that will be controlled with the C2. The first section of the .yml file has the label "all." Entries under this label should have the format `ip:interface` and EVERY host should be listed, regardless of whether it will be in other groups. The labels after that (for example "web" and "team1") will have entries formated as `ip` for each host that will be in that group. There's an example file located at `modules/daddy_tops/utils/example_input.yml`. Once the `hostbuilder` tool is used, hosts and groups should now be available/visible when using the `dt_client.py`, and actions can begin to be queued for hosts/groups.  

#### Using the CLI  
The CLI is `dt_client.py`, which is located in `clients/daddy_tops/`. When run, it will prompt the user for a username/password. Once entered, a token will be generated and used for all requests to the DaddyTops API. The prompt will look like the following:  

`DT>`  

From here there a lots of options. The `help` option will show the user the CLI options/arguments, `exit` will quit the session. [Located here](https://streamable.com/oj0zc) is a video that walks through some examples of how to use the tool.

