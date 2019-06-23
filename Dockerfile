# Derive the image from a golang image 
FROM golang:1.12.6-alpine

# Install git 1st
RUN apk add --update git

# Get the go-cards package
RUN go get github.com/wendysanarwanto/go-cards

# Build the go-cards package
RUN go build github.com/wendysanarwanto/go-cards

# Run the go-cards sample app
CMD ["go-cards"]
