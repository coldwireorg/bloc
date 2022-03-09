FROM golang:1.17.6 as builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /tmp/auth
COPY . .
RUN go build -a -o /tmp/auth/main .
FROM scratch
COPY --from=builder /tmp/auth/main /main
ENTRYPOINT ["/main"]