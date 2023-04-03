# Use Go 1.17-alpine as base image
FROM golang:1.17-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the module dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Install rsync and exclude unnecessary files
RUN apk add --no-cache rsync && \
    rsync -av --exclude=k8s --exclude=config/redis.conf --delete-excluded ./ /app && \
    go build -o app . && \
    find . -type f ! -name "app" -delete 

# Set the command to run the application
CMD ["/app/app"]