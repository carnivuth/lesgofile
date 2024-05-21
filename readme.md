# LESGOFILE
 
Simple project to transfer file over network 

The goal of the project is to have a simple daemon application to experiment deployment and distribution techniques

## INSTALLATION:

### FROM SOURCE

- install the make dependencies `go`

- clone repo

- run `sudo make install`

you can set `DESTDIR` and `prefix` variables in make to customize the installation directory

## USAGE

the main binary file can act as server and client, 

 - to run client:

 `lesgofile send "ip address" filename`

- to run server:
 `lesgofile recive`
