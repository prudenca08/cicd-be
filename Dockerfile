FROM golang:1.16-alpine3.14

WORKDIR /finalproject-be-recipe

COPY . .

RUN go mod download

RUN go build -o mainfile

EXPOSE 80/tcp

CMD ["./mainfile"]