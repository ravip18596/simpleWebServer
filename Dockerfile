FROM golang:1.10

WORKDIR $GOPATH/src
# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

RUN pwd  && echo $GOPATH
# Download all the dependencies
RUN cd src && go get -d -v ./...
# Install the package
RUN cd src && go install -v ./...
# running sever
RUN cd src && go build -o server

RUN pwd && ls

# This container exposes port 8080 to the outside world
EXPOSE 8000

# Run the executable
CMD ["./src/server"]