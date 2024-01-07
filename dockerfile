# Start with a base Go image
FROM golang:1.21.0-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application listens on
EXPOSE 8080

# Run the Go application
CMD ["./main"]