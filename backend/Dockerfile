FROM golang:1.23 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go install ./cmd/server

FROM gcr.io/distroless/static-debian12
COPY --from=build /go/bin/* /go/bin/
WORKDIR /app
COPY . .
CMD ["/go/bin/server"]
