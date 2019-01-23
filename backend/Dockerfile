ARG HTTP_PORT=3000

FROM golang:latest AS builder

WORKDIR /src

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

FROM scratch as final

COPY --from=builder /app /app

EXPOSE ${HTTP_PORT}

ENTRYPOINT [ "/app" ]