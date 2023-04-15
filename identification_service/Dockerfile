FROM golang:1.19-alpine

WORKDIR /app

# Install make
# Install dependencies
RUN apk update && apk add make


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ /app

RUN go build -o /identification_service

EXPOSE 8080

CMD ["make", "bin-deps"]
CMD ["make", "install"]
CMD [ "/identification_service" ]
