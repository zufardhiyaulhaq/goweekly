#################
# Base image
#################
FROM alpine:3.12 as goweekly-base

USER root

RUN addgroup -g 10001 goweekly && \
    adduser --disabled-password --system --gecos "" --home "/home/goweekly" --shell "/sbin/nologin" --uid 10001 goweekly && \
    mkdir -p "/home/goweekly" && \
    chown goweekly:0 /home/goweekly && \
    chmod g=u /home/goweekly && \
    chmod g=u /etc/passwd

ENV USER=goweekly
USER 10001
WORKDIR /home/goweekly

#################
# Builder image
#################
FROM golang:1.16-alpine AS goweekly-builder
RUN apk add --update --no-cache alpine-sdk
WORKDIR /app
COPY . .
RUN make build

#################
# Final image
#################
FROM goweekly-base

COPY --from=goweekly-builder /app/bin/goweekly /usr/local/bin

# Command to run the executable
ENTRYPOINT ["goweekly"]
