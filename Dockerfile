FROM golang:1.22-alpine
WORKDIR /app

COPY go.* /
RUN go mod download

COPY . .
RUN go build -o /webhealth-inspect 

EXPOSE 8080

ENTRYPOINT ["/webhealth-inspect"]
