FROM golang:1.19 as builder

WORKDIR /usr/src/fargate-sidecar-injector

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -v -o fargate-sidecar-injector ./pkg/cmd

FROM alpine:latest

WORKDIR /usr/src/fargate-sidecar-injector

COPY --from=builder /usr/src/proxy/fargate-sidecar-injector ./

EXPOSE 3003

ENTRYPOINT ["./fargate-sidecar-injector"]