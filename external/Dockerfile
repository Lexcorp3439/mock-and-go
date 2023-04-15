FROM golang:1.19-alpine

ARG DATABASE_DSN
RUN export DATABASE_DSN=$DATABASE_DSN

WORKDIR /app

# Install make
# Install dependencies
RUN apk update && apk add make


COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ /app

RUN go build -o /external

EXPOSE 8080

CMD ["make", "bin-deps"]
CMD [ "/external" ]
