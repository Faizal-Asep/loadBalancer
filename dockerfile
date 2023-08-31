FROM golang:1.21.0-alpine AS builder
# RUN apk add --no-cache ca-certificates
RUN apk --no-cache add tzdata
WORKDIR /go/src/microservice
COPY . /go/src/microservice
RUN CGO_ENABLED=0 go build -o loadbalancer *.go

FROM scratch
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/microservice/loadbalancer ./
COPY --from=builder /usr/share/zoneinfo/Asia/Jakarta /usr/share/zoneinfo/Asia/Jakarta
ENV TZ=Asia/Jakarta
CMD ["/loadbalancer"]