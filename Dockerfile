FROM golang:1.22-alpine

COPY . /go/src/lesgofile
WORKDIR /go/src/lesgofile

# get dependencies
RUN go mod tidy

# build
RUN go build -o /lesgofile
EXPOSE 50000

# run the server component
CMD ["/lesgofile", "serve"]
