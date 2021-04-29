FROM golang:1.16.3

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/w00k/pong

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8081 to the outside world
EXPOSE 8081

# Run the executable
CMD ["pong"]
