FROM golang:1.13-alpine AS builder

RUN apk --update add ca-certificates git gcc musl-dev make

RUN mkdir /builder

ADD . /builder

WORKDIR /builder

RUN go mod tidy

RUN go test ./...

RUN make build-linux

FROM scratch AS production

COPY --from=builder /builder/build/qr-generator /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

EXPOSE 3506

ENTRYPOINT ["/qr-generator"]