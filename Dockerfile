# BUILD CONTAINER
FROM golang:1.19.9-buster AS build

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

#first layers will be with dependencies, which change less often
COPY go.* /

#instructing Go tools that these are private repos; add -json for very verbose output
RUN go mod download

#next layer is the app which change often on localhost
COPY . /


RUN go build -o /app/server /cmd/clihands/main.go


# RUN CONTAINER
FROM alpine:3.11.11

COPY --from=build /app/server /app/server

#EXPOSE 8000
WORKDIR /app
#for Go profiling see PROFILING_PORT
#EXPOSE 8001

#ENTRYPOINT ["/app/server", "--port=8000","--host=0.0.0.0","--read-timeout=120m","--write-timeout=120m","--scheme=http"]
ENTRYPOINT ["/app/server"]
