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
COPY server.crt .
COPY server.key .

RUN go build -o /server ./cmd/server/*.go
RUN go build -o /database ./cmd/database/*.go

CMD ["/database"]

ENTRYPOINT [ "/server" ]
EXPOSE 8443
