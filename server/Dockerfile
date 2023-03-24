FROM golang:1.20.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /traffic-control

HEALTHCHECK --interval=60s --timeout=3s CMD curl -f http://localhost/health

CMD /traffic-control -x 10000 -y 10000 -z 2