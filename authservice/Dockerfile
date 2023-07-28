FROM golang:1.20 AS build-stage

WORKDIR /app

COPY . ./

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-authservice

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-gs-authservice /docker-gs-authservice


EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-authservice"]