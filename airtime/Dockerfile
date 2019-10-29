FROM golang

# All these steps will be cached
RUN mkdir /airtime
WORKDIR /airtime
COPY airtime/go.mod .
COPY airtime/go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY airtime/. .

RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher
