FROM golang:1.21

ENV CGO_ENABLED=0

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/ ./...

EXPOSE 8090

CMD ["art-gallery", "serve", "--http=0.0.0.0:8090"]
