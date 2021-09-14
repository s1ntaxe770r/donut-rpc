FROM golang:latest as builder

WORKDIR /app
COPY ./go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o donutservice

FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/donutservice /donutservice

EXPOSE 5050

ENTRYPOINT ["/donutservice"]
