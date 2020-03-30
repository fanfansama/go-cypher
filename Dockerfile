############################
# STEP 1 build executable binary
############################
# Start from the latest golang base image
FROM golang:latest AS builder

# Add Maintainer Info
LABEL François BERTHAULT "francois.berthault@talanlabs.com"

# Set the Current Working Directory inside the container
WORKDIR /app
ENV GOPATH=/app

#WORKDIR $GOPATH/src/fanfansama/go-cypher/
#COPY . .

RUN mkdir in out-enc out-dec
# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN ./get-packages.sh
# Using go get.
#RUN go get -d -v

# Build the Go app
# Build the binary.
RUN go build -o /app/main
#RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/main

############################
# STEP 2 build a small image
############################
FROM gcr.io/distroless/base

# Copy our static executable.
COPY --from=builder /app/main /app/main

COPY --from=builder /app/in /app/in
COPY --from=builder /app/out-enc /app/out-enc
COPY --from=builder /app/out-dec /app/out-dec


WORKDIR /app
# RUN mkdir in out-enc out-dec

# Expose port 4450 to the outside world
EXPOSE 4450

# Run the  binary.
CMD ["/app/main"]

# Command to run the executable
#CMD ["./main"]
#ENTRYPOINT ["/app/main"]

