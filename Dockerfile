# Start from golang base image
FROM golang:alpine as builder

#Add gcc
RUN apk add build-base

# Install git. (alpine image does not have git in it)
RUN apk update && apk add --no-cache git

# Set current working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all dependencies.
RUN go mod download

# Now, copy the source code
COPY . .

# Build the application.
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

FROM alpine

# Copy the Pre-built binary file
COPY --from=builder /app/bin/main .

# Run executable
CMD ["./main"]