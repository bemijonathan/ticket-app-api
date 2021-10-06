FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go get github.com/pilu/fresh

EXPOSE 4000

CMD ["fresh"]