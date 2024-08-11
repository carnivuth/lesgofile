# BRAINSTORM
## GOALS

quick installation and 0 configuration startup
per user configuration
change file transfer protocol to http? maybe be protocol agnostic
discovery?
cli interaction

## CLI STRUCTURE

lesgofile discover //will list discover servers in the area
lesgofile send  <file> <dest> // dest optional, if not provided, start discover and list discovered servers

## DISCOVERY
server will listen on udp port and reply with name:address to clients
clients will send broadcast discovery requests on the local network and get replies from servers

## TESTING
Use docker and docker compose to test network functionality
Develop dockerfile
