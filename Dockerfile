FROM golang:alpine as app-builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 go build -o dice cmd/server/server.go

FROM scratch
COPY --from=app-builder /go/src/app/dice /dice
ENTRYPOINT ["/dice"]
