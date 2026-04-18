FROM golang:1.22-bookworm AS builder
WORKDIR /app
COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal
RUN go build -o /subtitle ./cmd/subtitle

FROM python:3.11-slim-bookworm
WORKDIR /app

RUN apt-get update \
    && apt-get install -y --no-install-recommends ffmpeg \
    && rm -rf /var/lib/apt/lists/*

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

COPY --from=builder /subtitle /usr/local/bin/subtitle
COPY scripts ./scripts

ENTRYPOINT ["subtitle"]
CMD ["-h"]
