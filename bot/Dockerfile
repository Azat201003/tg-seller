FROM golang:alpine

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o app .

CMD [ "./app" ]
