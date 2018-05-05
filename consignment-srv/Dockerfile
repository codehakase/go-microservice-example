FROM golang:1.9.0 as builder

# Set our workdir to our current service in the gopath
WORKDIR /home/codehakase/gocode/src/github.com/codehakase/go-microservice-example/consignment-srv

# Copy the current code into our workdir
COPY . .

# Here we're pulling in godep, which is a dependency manager tool.
RUN go get -u github.com/golang/dep/cmd/dep

# Create a dep project, and run `ensure`, which will pull in all
# of the dependencies within this directory.
RUN dep init && dep ensure

# Build the binary, with a few flags which will allow
# us to run this binary in Alpine/Debian.
RUN CGO_ENABLED=0 GOOS=linux go build -o consignment-srv -a -installsuffix cgo main.go repository.go handler.go datastore.go


# Here we're using a second FROM statement, which is strange,
# but this tells Docker to start a new build process with this
# image.
FROM debian:latest

# Security related package, good to have.
RUN apk --no-cache add ca-certificates

# Same as before, create a directory for our app.
RUN mkdir /app
WORKDIR /app

COPY --from=builder /home/codehakase/go/src/github.com/codehakase/go-microservice-example/consignment-srv/consignment-srv .

CMD ["./consignment-srv"]
