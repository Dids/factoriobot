# Compiler image
FROM didstopia/base:go-alpine-3.5 AS go-builder

# Copy the project 
COPY . /tmp/factoriobot/
WORKDIR /tmp/factoriobot/

# Install dependencies
RUN make deps

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/factoriobot



# Runtime image
FROM scratch

# Copy certificates
COPY --from=go-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy the built binary
COPY --from=go-builder /go/bin/factoriobot /go/bin/factoriobot

# Expose environment variables
ENV DISCORD_BOT_TOKEN       ""
ENV DISCORD_BOT_CHANNEL_ID  ""
ENV RCON_HOST               "localhost"
ENV RCON_PORT               "34198"
ENV RCON_PASSWORD           ""

# Run the binary
ENTRYPOINT ["/go/bin/factoriobot"]
