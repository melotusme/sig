FROM golang:1-alpine as builder

RUN apk --no-cache --no-progress add openssl git ca-certificates tzdata make \
&& update-ca-certificates \
&& rm -rf /var/cache/apk/*

WORKDIR /go/kws_sig

COPY go.mod .
COPY go.sum .
RUN GO111MODULE=on GOPROXY=https://proxy.golang.org go mod download
COPY . .
RUN make build

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/kws_sig/kws_sig .
COPY --from=builder /go/kws_sig/keys/ca.key .

ENTRYPOINT ["/kws_sig"]