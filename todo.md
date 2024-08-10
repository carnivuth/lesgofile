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

discover structure:

servers open http socket
clients send multicast requests to specific http socket
