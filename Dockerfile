# Use the latest Go version (1.24.1)
FROM golang:1.24.1-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy all files to the container
COPY . .

# Download dependencies
RUN go mod tidy

# Build the application
RUN go build -o app .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./app"]
