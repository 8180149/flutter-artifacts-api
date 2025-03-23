# Use the official Golang image
FROM golang:1.21-alpine

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
