# Force an amd64 Ubuntu base on any host.
FROM --platform=linux/amd64 ubuntu:22.04 AS builder

WORKDIR /app
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && \
    apt-get install -y --no-install-recommends \
      build-essential \
      wget \
      ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Install Go 1.24
ENV GOVER="1.24.0"
RUN wget https://go.dev/dl/go${GOVER}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GOVER}.linux-amd64.tar.gz && \
    rm go${GOVER}.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"
ENV CGO_ENABLED=1

COPY . .

RUN make clean && make

FROM ubuntu:22.04
WORKDIR /app
COPY --from=builder /app/blackcat .
COPY --from=builder /app/plugins/*.so plugins/
EXPOSE 8080
CMD ["./blackcat"]
