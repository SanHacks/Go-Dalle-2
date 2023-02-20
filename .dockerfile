# Use an official Go runtime as a parent image
FROM golang:1.16-alpine3.13

# Install MySQL and dependencies
RUN apk update && \
    apk add --no-cache mysql mysql-client

# Set environment variables
ENV MYSQL_DATABASE=root \
    MYSQL_ROOT_PASSWORD=root \
    MYSQL_USER=root \
    MYSQL_PASSWORD=root

# Copy the Go application code into the container
COPY . /app

# Set the working directory to the Go application
WORKDIR /app


# Build the Go application
RUN go build -o main .

# Expose port 8080 for Google App Engine
EXPOSE 8080

# Start the Go application and MySQL
CMD ["./main"]
