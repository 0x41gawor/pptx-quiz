FROM golang:1.24.1-alpine

WORKDIR /pptx-quiz

COPY go.mod ./
COPY go.sum ./
COPY ./cmd ./cmd
COPY ./internal ./internal
COPY Makefile ./
COPY data.csv ./

RUN go mod download
RUN apk add --no-cache make
RUN make build

EXPOSE 3456

CMD ["./bin/pptx-quiz"]