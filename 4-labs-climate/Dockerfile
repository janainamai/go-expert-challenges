FROM golang:1.22.3 AS build
WORKDIR /app
COPY . .
WORKDIR /app/cmd/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/cloudrun

FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
ENTRYPOINT ["./cloudrun"]