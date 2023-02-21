# Use an official Go runtime as a parent image
FROM golang:1.19-alpine3.13

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


FROM golang:1.17

# Install any additional dependencies your app requires
RUN apt-get update && apt-get install -y \
    libmysqlclient-dev \
    && rm -rf /var/lib/apt/lists/* \
    && go get github.com/go-sql-driver/mysql \
    && go get github.com/gorilla/mux \
    && go get github.com/gorilla/handlers \
    && go get github.com/gorilla/context \
    && go get github.com/gorilla/sessions \
    && go get github.com/gorilla/securecookie \

# Copy your app's source code to the container
COPY . /app

# Set the working directory
WORKDIR /app

# Build the app
RUN go build -o ./bin/app

# Set the entrypoint for the container
ENTRYPOINT ["/app/bin/app"]
