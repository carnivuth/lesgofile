# LESGOFILE

Simple project to transfer file over network

The goal of the project is to have a simple daemon application to experiment deployment and distribution techniques and some go programming

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
