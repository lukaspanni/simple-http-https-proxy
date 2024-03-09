FROM golang:bookworm AS builder

WORKDIR /build
COPY . .
RUN ./build.sh


FROM alpine:latest AS runner
COPY --from=builder /build/proxy .
ENTRYPOINT ["./proxy", "--source",":80"]