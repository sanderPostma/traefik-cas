FROM golang:1.14-alpine as build-env

# Setup
RUN mkdir -p /go/src/traefik-cas
WORKDIR /go/src/traefik-cas

# Add libraries
RUN apk add --no-cache git

# Copy & build
ADD . /go/src/traefik-cas/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -installsuffix nocgo -o /traefik-cas

# Copy into scratch container
FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /traefik-cas ./
ENTRYPOINT ["./traefik-cas"]
