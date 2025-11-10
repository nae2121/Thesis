FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/api ./cmd/api


FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /bin/api /app/api
COPY --from=build /src /app
EXPOSE 8080
ENTRYPOINT ["/app/api"]