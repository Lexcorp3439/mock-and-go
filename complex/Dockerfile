FROM golang:1.19-alpine

ARG DATABASE_DSN
ENV DATABASE_DSN=${DATABASE_DSN}

WORKDIR /app

# Install make
# Install dependencies
RUN apk update && apk add make

RUN apk add --virtual build-dependencies
RUN apk add --no-cache build-base gcc

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ /app

RUN go build -o /complex

EXPOSE 8080

CMD ["make", "bin-deps"]
CMD ["make", "db-up"]
CMD [ "/complex" ]
