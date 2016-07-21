## Kaffebot

#### Commands
* IP-STATE
    * Tells all the ips associated with all interfaces
* KAFFE-STATE
    * You know what
    
#### Developing
Install Go (https://golang.org). Clone the repository and set GOPATH env variable to the root of the repo.

    
#### Deploying new version
Get the ip address from the bot (or nmap if it is down).
Currently bot is running in a `screen`, retach the screen `screen -dr` and ctrl-C to stop the bot. 
Build new version of the bot (`GOARCH=arm go install kaffebot`) on local machine
and copy it to the raspi (`scp bin/linux_arm/kaffebot pi@IP_ADDR:~/`).
Start the bot in the screen `~/kaffebot` and `ctrl-a d` detaches the screen.


#### TODO
init scripts to start automagically at the raspi