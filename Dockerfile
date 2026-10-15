FROM golang:1.25-alpine AS build

ARG TARGETOS=linux
ARG TARGETARCH=amd64

RUN apk add --no-cache ca-certificates git tzdata

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o user-service .

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=build /app/user-service /app/user-service
COPY --from=build /app/config /app/config
COPY --from=build /app/docs /app/docs
COPY --from=build /app/pkg /app/pkg

EXPOSE 5001 6001 7001 8001

USER nonroot:nonroot
ENTRYPOINT ["/app/user-service"]
CMD ["start"]
