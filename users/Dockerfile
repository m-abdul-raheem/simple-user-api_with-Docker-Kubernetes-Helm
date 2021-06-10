FROM golang:1.15.6-alpine

WORKDIR /workdir/app/

COPY go.mod .

RUN go mod download
RUN go mod verify

COPY . .

RUN go build -o users ./code

FROM alpine
COPY --from=0 /workdir/app/users .
ENTRYPOINT ["./users"]