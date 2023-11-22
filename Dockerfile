FROM golang:1.16-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum main.go ./
RUN go mod vendor && go build

FROM debian:bullseye-slim AS runner

RUN apt update && apt install wget -y
RUN wget -O PPL.deb https://dl4jz3rbrsfum.cloudfront.net/software/PPL_64bit_v1.4.1.deb && dpkg -i *.deb

COPY --from=builder /app/pwrstat-exporter /app/pwrstat-exporter

COPY init.sh /app
RUN chmod +x /app/init.sh

ENTRYPOINT "/app/init.sh"