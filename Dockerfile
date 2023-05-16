# Build stage
FROM golang:alpine AS build
RUN apk add --no-cache gcc g++ git openssh-client

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . ./

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /freshpoint

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir -p /storage

COPY --from=build /newsletter /newsletter

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

CMD ["/newsletter"]