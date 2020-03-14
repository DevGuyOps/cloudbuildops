# Build Container
FROM golang:alpine as builder
WORKDIR /go/src/app
COPY . .

RUN go build -o cloudbuildops cmd/cloudbuildops/*.go

# Run Container
FROM alpine
COPY --from=builder /go/src/app/cloudbuildops /
RUN chmod +x /cloudbuildops

ENTRYPOINT ["./cloudbuildops"]
