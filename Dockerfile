FROM golang:1.19

WORKDIR /app

RUN go mod init teste-soma

COPY . .

RUN go build -o soma

CMD ["./soma"]