FROM golang:1.22.3 AS build
WORKDIR /app
COPY . .
WORKDIR /app/cmd/app
RUN apt-get update && apt-get install -y ca-certificates
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/cloudrun

FROM scratch
WORKDIR /app
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/cloudrun .
ENTRYPOINT ["./cloudrun"]