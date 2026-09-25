FROM golang:1.27.1-alpine3.24 AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o /build/app main.go

FROM alpine:3.24 AS final

ARG MIGRATIONS_DIR_PATH
ARG API_WORKDIR
ARG API_BIN_PATH

WORKDIR ${API_WORKDIR}

COPY --from=builder /build/db/migrations ${MIGRATIONS_DIR_PATH}
COPY --from=builder /build/app ${API_BIN_PATH}/app
CMD ["app"]