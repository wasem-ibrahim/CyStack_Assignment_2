# CyStack_Assignment_2
Code built by Go language to build a client-server application that collects basic user's information about any OSX machine.

## Features:
### The application collects the following data: 
* Host name
* IP Address
* Memory Utilization
* Disk Utilization
* All Local Users in OSX
* All Running Process List for OSX 
* All installed Application

### The application then collects the information and stores them inside of a "SQLite" database

## Note: 
The only needed thing to be changed is the "path" that leads to the ".db" file.
The line that needs to changed is inside of the "SendData" file which is inside of the "database" file inside of the "client" file.

## Steps to run build the damage file (As it was relatively big to be uploaded on github):
1. Create an "exec" file using the following command in terminal `go build -o client`.
2. Move the file that was built inside of the "dmg" file
3. Change the "path" inside of the ".plist" file to your path.
4. run the following command `hdiutil create -srcfolder path/to/directory -format UDZO client.dmg` replacing "path/to/directory" with the actual path at your device.



