# Video Subtitle Pipeline (Go + Whisper + LibreTranslate)

A simple backend project that processes a local video and produces a subtitle-burned output video.

Pipeline steps:
1. Extract audio from input video with **FFmpeg**
2. Transcribe audio with **OpenAI Whisper** (Python script)
3. Generate **SRT** subtitles
4. Translate SRT lines using **LibreTranslate-compatible API**
5. Burn translated subtitles into the output video with **FFmpeg**

## Architecture

The project follows a lightweight clean architecture split:

- `cmd/subtitle`: CLI entrypoint and wiring
- `internal/usecase`: pipeline orchestration
- `internal/ports`: service interfaces (video/transcription/translation/SRT)
- `internal/infrastructure`: concrete implementations
  - `video`: FFmpeg integration
  - `transcription`: Whisper script wrapper
  - `translation`: LibreTranslate script wrapper
  - `srt`: SRT writer
- `scripts/`: Python scripts used by Go services

## Prerequisites (local run)

- Go 1.22+
- Python 3.11+
- FFmpeg installed and available in PATH

Python dependencies:

```bash
pip install -r requirements.txt
```

## CLI usage

```bash
go run ./cmd/subtitle \
  -input ./samples/input.mp4 \
  -output ./samples/output_es.mp4 \
  -source-lang auto \
  -target-lang es
```

Available flags:

- `-input` (required): input video path
- `-output`: output video path (default `output_subtitled.mp4`)
- `-source-lang`: source language code for translation (default `auto`)
- `-target-lang`: target language code (default `es`)
- `-workdir`: temp working directory (default uses system temp)
- `-keep-artifacts`: keep intermediate WAV and SRT files
- `-whisper-model`: Whisper model name (default `base`)
- `-translate-endpoint`: translation API endpoint (default `https://libretranslate.com/translate`)

## Docker

Build:

```bash
docker build -t subtitle-pipeline .
```

Run (single command):

```bash
docker run --rm \
  -v "$PWD:/data" \
  subtitle-pipeline \
  -input /data/input.mp4 \
  -output /data/output_es.mp4 \
  -target-lang es
```

> Note: Whisper downloads model weights on first run, which can take time.

## Notes

- Translation is line-based on subtitle text lines to keep implementation simple.
- You can replace `-translate-endpoint` with your own LibreTranslate-compatible service.
