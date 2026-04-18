#!/usr/bin/env python3
import argparse
import json

import whisper


def main() -> None:
    parser = argparse.ArgumentParser(description="Transcribe audio with OpenAI Whisper")
    parser.add_argument("--audio", required=True, help="Path to audio file")
    parser.add_argument("--output", required=True, help="Path to JSON output")
    parser.add_argument("--model", default="base", help="Whisper model size")
    args = parser.parse_args()

    model = whisper.load_model(args.model)
    result = model.transcribe(args.audio)

    segments = [
        {
            "start": segment["start"],
            "end": segment["end"],
            "text": segment["text"].strip(),
        }
        for segment in result.get("segments", [])
    ]

    with open(args.output, "w", encoding="utf-8") as f:
        json.dump(segments, f, ensure_ascii=False)


if __name__ == "__main__":
    main()
