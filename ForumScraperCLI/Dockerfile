# Use the official Go image
FROM golang:1.23 AS builder

# Set up working directory
WORKDIR /app


# Copy application files
COPY . .

# Copy Go module files and download dependencies
COPY go.mod go.sum ./
COPY static static
RUN go mod download

# Build the Go application
RUN go build -o forum_monitoring_tools .

# Use a lightweight Debian image
FROM debian:bookworm-slim

# Create a volume for screenshots
WORKDIR /app

# Set up working directory
COPY screenshots screenshots
COPY static static

# Install dependencies
RUN apt-get update && \
    apt-get install -y tor wget gnupg2 curl && \
    curl -sSL https://dl.google.com/linux/linux_signing_key.pub | gpg --dearmor -o /usr/share/keyrings/google-chrome.gpg && \
    echo "deb [arch=amd64 signed-by=/usr/share/keyrings/google-chrome.gpg] http://dl.google.com/linux/chrome/deb/ stable main" > /etc/apt/sources.list.d/google-chrome.list && \
    apt-get update && \
    apt-get install -y google-chrome-stable && \
    rm -rf /var/lib/apt/lists/*

# Set up Tor configuration
RUN echo "SOCKSPort 0.0.0.0:9150" >> /etc/tor/torrc

# Copy the Go application from the builder stage
COPY --from=builder /app/forum_monitoring_tools /usr/local/bin/forum_monitoring_tools

# Command to start Tor and then run the application
CMD ["sh", "-c", "service tor start && exec /usr/local/bin/forum_monitoring_tools"]
