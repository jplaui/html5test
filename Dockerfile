FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/jplaui/html5test

# Build the outyet command inside the container.
# RUN go get....
RUN go install github.com/jplaui/html5test

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/html5test

# Document that the service listens on port 8080.
EXPOSE 8080

