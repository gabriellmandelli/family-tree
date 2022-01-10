FROM golang:1.15-alpine as build

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o family-tree ./cmd/family-tree/.

FROM alpine
RUN apk --no-cache add tzdata
WORKDIR /app
USER 1000
COPY --from=build /build/family-tree .
ADD build/package/docker/service/entrypoint.sh /

ENTRYPOINT [ "/entrypoint.sh" ]