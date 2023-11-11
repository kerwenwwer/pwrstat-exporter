FROM golang:1.16-bullseye

WORKDIR /app

RUN wget -O PPL.deb https://dl4jz3rbrsfum.cloudfront.net/software/PPL_64bit_v1.4.1.deb && dpkg -i *.deb

COPY go.mod go.sum main.go ./
RUN go mod vendor && go build

COPY init.sh ./
RUN chmod +x /app/init.sh

ENTRYPOINT "/app/init.sh"