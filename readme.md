# LESGOFILE
 
just a simple project to transfer file over network 

## installation guide:
- from source code
    - install go compiler
    - clone repo
    - launch `deploy.sh GOOS GOARCH filename` 
    - extract the .tar.gz file
    - set correct variables in `settings.json` file
    - enjoy :)

## parameters
 - run client:
 `lesgofile send "ip address" filename`
- run server:
 `lesgofile recive`
