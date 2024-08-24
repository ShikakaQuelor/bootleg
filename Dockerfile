FROM golang:1.22.5-alpine AS golang

WORKDIR /bootleg
COPY . .

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /server ./cmd/ws/

FROM gcr.io/distroless/static-debian11

COPY --from=golang /server .
COPY static ./static

CMD ["/server"]
