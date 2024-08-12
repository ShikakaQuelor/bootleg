FROM golang:1.22.5-alpine as golang

WORKDIR /bootleg
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bootleg .

FROM gcr.io/distroless/static-debian11

COPY --from=golang /bootleg .

EXPOSE 8080

CMD ["/bootleg"]
