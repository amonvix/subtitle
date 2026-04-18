#!/usr/bin/env python3
import argparse
import re

import requests

TIME_LINE_PATTERN = re.compile(r"\d{2}:\d{2}:\d{2},\d{3}\s+-->\s+\d{2}:\d{2}:\d{2},\d{3}")


def should_translate(line: str) -> bool:
    stripped = line.strip()
    if not stripped:
        return False
    if stripped.isdigit():
        return False
    if TIME_LINE_PATTERN.fullmatch(stripped):
        return False
    return True


def translate_text(text: str, endpoint: str, source: str, target: str) -> str:
    payload = {
        "q": text,
        "source": source,
        "target": target,
        "format": "text",
    }
    response = requests.post(endpoint, json=payload, timeout=30)
    response.raise_for_status()
    data = response.json()
    return data.get("translatedText", text)


def main() -> None:
    parser = argparse.ArgumentParser(description="Translate SRT using LibreTranslate-compatible API")
    parser.add_argument("--input", required=True, help="Input SRT path")
    parser.add_argument("--output", required=True, help="Output SRT path")
    parser.add_argument("--source", default="auto", help="Source language code")
    parser.add_argument("--target", required=True, help="Target language code")
    parser.add_argument("--endpoint", required=True, help="Translation endpoint URL")
    args = parser.parse_args()

    with open(args.input, "r", encoding="utf-8") as f:
        lines = f.readlines()

    output_lines = []
    for line in lines:
        if should_translate(line):
            translated = translate_text(line.strip(), args.endpoint, args.source, args.target)
            output_lines.append(translated + "\n")
        else:
            output_lines.append(line)

    with open(args.output, "w", encoding="utf-8") as f:
        f.writelines(output_lines)


if __name__ == "__main__":
    main()
