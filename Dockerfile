#
# SETUP
#
FROM docker.io/golang:1.19-alpine AS builder

RUN apk add --update --no-cache ca-certificates git just binutils
WORKDIR /build

# Alpine compat and general irritability
ENV CGO_ENABLED=0

#
# CACHE
#

COPY go.mod go.sum justfile .
RUN just precache

#
# BUILD
#

COPY . .
RUN just release

#
# CLEAN
#

FROM scratch
COPY --from=builder /build/example-ent-crdb /app
ENTRYPOINT ["/app/example-ent-crdb"]
