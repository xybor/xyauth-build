FROM golang:1.18-alpine

WORKDIR /xyauth

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY configs configs/
COPY cmd cmd/
COPY internal internal/
COPY pkg pkg/
COPY web web/
COPY server.crt server.crt
COPY server.key server.key
COPY server.crt server.crt
COPY server.key server.key

RUN go build -o /server ./cmd/server/*.go
RUN go build -o /database ./cmd/database/*.go

CMD ["/database", "migrate"]

ENTRYPOINT [ "/server" ]
EXPOSE 8443
