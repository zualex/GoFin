# Start from golang base image
FROM golang:alpine as base

FROM base as dev

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base

# Setup folders
RUN mkdir /app
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .
# COPY .env .

# # Download all the dependencies
RUN go get -d -v ./...

# # Install the package
RUN go install -v ./...

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go install github.com/cosmtrek/air@latest

RUN go mod tidy

CMD ["/app/start.sh"]

# # Build the Go app
# RUN go build -o /build

# # Expose port 8080 to the outside world
# EXPOSE 8080

# # Run the executable
# CMD [ "/build" ]