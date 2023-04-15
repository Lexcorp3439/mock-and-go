FROM golang:1.19-alpine

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

RUN go build -o /users_service

EXPOSE 8080

ARG DATABASE_DSN
ENV DATABASE_DSN=$DATABASE_DSN

RUN make bin-deps
#RUN make db-up
CMD [ "/users_service" ]
