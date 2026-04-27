# Stage 1: BUILD
FROM --platform=$BUILDPLATFORM golang:1.26.1 AS builder

# mod caching 
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

# copying surce
COPY . .

# cross-compile respecting target platform
ARG TARGETOS
ARG TARGETARCH

# building the app
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o app ./cmd/solar-frontend

# Stage 2: final image
FROM --platform=$BUILDPLATFORM alpine:3.23.3

# copying only the app in the final image
WORKDIR /app
COPY --from=builder /build/app .

# copying static files
COPY web ./web

RUN apk add --no-cache tzdata

ENV TZ=Europe/Rome

ENV ENERGY_BASE_URL="http://localhost:80"

EXPOSE 8080

CMD ["./app"]