FROM golang:latest as build

WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download the Go module dependencies
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /myapp main.go
 
FROM alpine:latest as run

# Copy the application executable from the build image
COPY --from=build /myapp /myapp

WORKDIR /app
EXPOSE 3000
CMD ["/myapp"]
   