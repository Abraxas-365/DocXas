FROM golang:alpine

# Add Maintainer info
LABEL maintainer="Luis F. Miranda"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /documentator
WORKDIR /documentator

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all the dependencies
# RUN go get -d -v ./...

# Install the package
RUN go mod tidy

# Build the Go app
RUN go build cmd/main.go

# Expose port 3000 to the outside world
EXPOSE 3000:3000

# Run the executable
CMD ["./main"]
