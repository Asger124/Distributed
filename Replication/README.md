# Mandatory hand-in 5 Auction service

This auction service uses a client-server archtecture and implements Active-replication in an Auction service. 
In order to run the program you need to start three servers and at least one client.
Servers will run on port 5000 + the id you enter in the terminal. Clients will run on port 6000 + the id.
All actions in the system will be logged in the log.txt file. The system is resilient to 2/3 server nodes crashing.

**How to start servers**
- Delete contents of log.txt file or the delete the file itself
- open 3 terminals
- Navigate to Replication/Server folder
- In each of the three terminals run one of the commands below
- type in following command: go run . -id 1
- type in following command: go run . -id 2
- type in following command: go run . -id 3

Note: servers cannot have the same id

**How to run clients**
- Ensure that three Servers are running
- open a new terminal
- Navigate to Replication/Client folder 
- type in the following command: go run . -id 1
- You can run as many clients as you want - using different terminals, just increment the id with each new client

The Auction will run for one minute. You can change this at line 20 in server_api.go