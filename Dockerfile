# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL François BERTHAULT "francois.berthault@talanlabs.com"

# Set the Current Working Directory inside the container
WORKDIR /app

ENV GOPATH=/app

RUN mkdir in out-enc out-dec
# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN ./get-packages.sh

# Build the Go app
RUN go build -o main .

# Expose port 4450 to the outside world
EXPOSE 4450

# Command to run the executable
CMD ["./main"]