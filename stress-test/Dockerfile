FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o st ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/st .
ENTRYPOINT ["./st"]