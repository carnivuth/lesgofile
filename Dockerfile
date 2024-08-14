FROM golang:1.22-alpine

# coping default config
COPY etc/lesgofile.json /etc/lesgofile/lesgofile.json

COPY . /go/src/lesgofile
WORKDIR /go/src/lesgofile

# get dependencies
RUN go mod tidy

# build
RUN go build -o /lesgofile
EXPOSE 50000

CMD ["/lesgofile", "serve"]
