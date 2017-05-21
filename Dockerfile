FROM ubuntu:16.04
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src
ENV GOBIN=/go/bin

# Build the outyet command inside the container.
# RUN go get....
RUN go install /go/src/server.go

# Document that the service listens on port 8080.
EXPOSE 8080

CMD ["/go/bin/server"]

