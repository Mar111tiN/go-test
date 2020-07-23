FROM golang:buster AS buildstage

RUN apt-get update --fix-missing && \
    apt-get install git && \
    apt-get clean

WORKDIR $GOPATH/src/go-test

# copy code (or use git to fetch the code?)
COPY . .

# fetch dependencies
RUN go mod download
RUN go mod verify

# build the binary
#  GOBIN is already set to /go/bin
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/go-test

FROM scratch 

COPY --from=buildstage /go/bin/go-test /bin/go-test

ENTRYPOINT ["/bin/go-test"]