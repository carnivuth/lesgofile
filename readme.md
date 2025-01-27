# LESGOFILE

Just a simple project to transfer file over network, the scope is to practice with deployment techniques, docker, go low level networking applications and system,

## FEATURES

- server and client included in a single binary
- discovery features
- multiple deployment strategies as packages for all modern distros and docker
- configuration from file

## INSTALLATION:

downlad latest release from the relase page

### FROM SOURCE

- install the make dependencies `go`

- clone repo

- run `go build`


## USAGE

the main binary file can act as server and client,

to run client:

```bash
lesgofile send <address> filename
```

you can also pipe the file name to the lesgofile client:

```bash
echo "filename" | lesgofile send <address>
```

to run server:

```bash
lesgofile serve
```
to search for servers in lan:

```bash
lesgofile discover
```
