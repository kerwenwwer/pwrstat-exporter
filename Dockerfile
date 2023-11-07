FROM golang:1.16-bullseye

WORKDIR /app

COPY *.deb ./
RUN dpkg -i *.deb

COPY go.mod go.sum main.go ./
RUN go mod vendor
RUN go build

COPY init.sh ./
RUN chmod +x /app/init.sh

ENTRYPOINT "/app/init.sh"