FROM golang:1.20 AS build-stage

WORKDIR /app

COPY . ./

RUN go mod download


RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-relservice

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-gs-relservice /docker-gs-relservice


EXPOSE 8082

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-relservice"]